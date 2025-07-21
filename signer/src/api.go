package signer

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/girvel/signer/docs"
)

// TYPES //

type Cryptographer interface {
    Sign(data string) ([]byte, error)
    Verify(data string, signature []byte) error
    Public() string
}

type deps struct {
    cryptographer Cryptographer
}

// ENDPOINTS //

type ErrorResponse struct {
    Error string `json:"error"`
}

// @Summary Get the public key
// @Description Get the public key used for signature verification
// @Produce plain
// @Success 200 {string} string "The key"
// @Router /public [get]
func (d *deps) public(c *gin.Context) {
    c.String(http.StatusOK, d.cryptographer.Public())
}

type SignBody struct {
    Text string `json:"text"`
}

type SignOkResponse struct {
    DatedText string `json:"dated_text"`
	Signature string `json:"signature"`
}

// @Summary Sign the given text
// @Description Get an RSA PSS signature for given text
// @Accept json
// @Produce json
// @Param body body signer.SignBody true "Text to sign"
// @Success 200 {object} signer.SignOkResponse "Returns the signature"
// @Failure 400 {object} signer.ErrorResponse "Can't bind JSON from body"
// @Failure 500 {object} signer.ErrorResponse "Issues with cryptography algorithm"
// @Router /sign [post]
func (d *deps) sign(c *gin.Context) {
    var body SignBody
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
        return
    }
    
	// TODO author env var
	datedText := fmt.Sprintf(
		"%s\n\nSigned %s by %s",
		body.Text,
		time.Now().Format("2006-01-02 15:04:05"),
		"girvel",
	)
    signature, err := d.cryptographer.Sign(datedText)

    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
        return
    }

	c.JSON(http.StatusOK, SignOkResponse{
		DatedText: datedText, 
		Signature: base64.StdEncoding.EncodeToString(signature),
	})
}

type VerifyBody struct {
    Data string `json:"data"`
    Signature []byte `json:"signature"`
}

type VerifyOkResponse struct {
    Correct bool `json:"correct"`
    Message string `json:"message"`
}

// @Summary Verify text + signature
// @Description Verify that given signature matches given text, using the public key
// @Accept json
// @Produce json
// @Param body body signer.VerifyBody true "Data-signature pair to verify; signature can be passed as base-64 string"
// @Success 200 {object} signer.VerifyOkResponse
// @Failure 400 {object} signer.ErrorResponse
// @Router /verify [post]
func (d *deps) verify(c *gin.Context) {
    var body VerifyBody
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
        return
    }

    err := d.cryptographer.Verify(body.Data, body.Signature)

    result := VerifyOkResponse{
        Correct: err == nil,
    }

    if err != nil {
        result.Message = err.Error()
    }

    c.JSON(http.StatusOK, result)
}

// FACTORY //

func CreateAPI(cryptographer Cryptographer) *gin.Engine {
    d := &deps{cryptographer}
    g := gin.Default()
    
    g.GET("/public", d.public)
    g.POST("/sign", d.sign)
    g.POST("/verify", d.verify)

    g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return g
}
