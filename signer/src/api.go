package signer

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/girvel/signer/signer/docs"
)

// TYPES //

type Cryptographer interface {
    Sign(data string) ([]byte, error)
    Verify(data string, signature []byte) error
    Public() string
}

type deps struct {
    Cryptographer
	*Env
}

// ENDPOINTS //

type ErrorResponse struct {
    Error string `json:"error"`
}

type SignaturePair struct {
    DatedText string `json:"dated_text"`
	Signature string `json:"signature"`
}

// @Summary Get the public key
// @Description Get the public key used for signature verification
// @Produce plain
// @Success 200 {string} string "The key"
// @Router /public [get]
func (d *deps) public(c *gin.Context) {
    c.String(http.StatusOK, d.Cryptographer.Public())
}

type SignBody struct {
    Text string `json:"text"`
}

// @Summary Sign the given text
// @Description Get an RSA PSS signature for given text
// @Accept json
// @Produce json
// @Param body body signer.SignBody true "Text to sign"
// @Success 200 {object} signer.SignaturePair "Returns the signature"
// @Failure 400 {object} signer.ErrorResponse "Can't bind JSON from body"
// @Failure 500 {object} signer.ErrorResponse "Issues with cryptography algorithm"
// @Router /sign [post]
func (d *deps) sign(c *gin.Context) {
    var body SignBody
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
        return
    }
    
	datedText := fmt.Sprintf(
		"%s\n\nSigned %s by %s with public key %s/public",
		body.Text,
		time.Now().Format("2006-01-02 15:04:05 MST"),
		d.Env.HolderName,
		d.Env.ServiceAddress,
	)
    signature, err := d.Cryptographer.Sign(datedText)

    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
        return
    }

	c.JSON(http.StatusOK, SignaturePair{
		DatedText: datedText, 
		Signature: base64.StdEncoding.EncodeToString(signature),
	})
}

type VerifyOkResponse struct {}

// TODO either endpoint functions are public too, or types are private

// @Summary Verify text + signature
// @Description Verify that given signature matches given text; quality of life feature, can be done locally with the public key.
// @Accept json
// @Produce json
// @Param body body signer.SignaturePair true "Data-signature pair to verify; signature can be passed as base-64 string"
// @Success 200 {object} signer.VerifyOkResponse "Signature matches"
// @Failure 400 {object} signer.ErrorResponse "Bad JSON"
// @Failure 409 {object} signer.ErrorResponse "Signature doesn't match"
// @Router /verify [post]
func (d *deps) verify(c *gin.Context) {
    var body SignaturePair
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
        return
    }

	rawSignature, err := base64.StdEncoding.DecodeString(body.Signature)
	if err != nil {
	    c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

    err = d.Cryptographer.Verify(body.DatedText, rawSignature)

    if err != nil {
		c.JSON(http.StatusConflict, ErrorResponse{err.Error()})
		return
    }

    c.JSON(http.StatusOK, VerifyOkResponse{})
}

// FACTORY //

func CreateAPI(cryptographer Cryptographer, env *Env) *gin.Engine {
    d := &deps{cryptographer, env}
    g := gin.Default()
    
    g.GET("/public", d.public)
    g.POST("/sign", d.sign)
    g.POST("/verify", d.verify)

    g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return g
}
