package signer

import (
	"encoding/base64"
	"net/http"

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

// @Summary Get the public key
// @Description Get the public key used for signature verification
// @Produce plain
// @Success 200 {string} string "The key"
// @Router /public [get]
func (d *deps) public(c *gin.Context) {
    c.String(http.StatusOK, d.cryptographer.Public())
}

// @Summary Sign the given text
// @Description Get an RSA PSS signature for given text
// @Accept plain
// @Produce plain
// @Param body body string true "Text to sign"
// @Success 200 {string} string "Returns the signature"
// @Failure 415 {string} string "Content-Type is not text/plain"
// @Failure 400 {string} string "Can't read body"
// @Failure 500 {string} string "Issues with cryptography algorithm"
// @Router /sign [post]
func (d *deps) sign(c *gin.Context) {
    if c.GetHeader("Content-Type") != "text/plain" {
        c.String(http.StatusUnsupportedMediaType, "Content-Type must be text/plain")
        return
    }

    body, err := c.GetRawData()
    if err != nil {
        c.String(http.StatusBadRequest, "Error reading body")
        return
    }
    
    signature, err := d.cryptographer.Sign(string(body))

    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }

    c.String(http.StatusOK, base64.StdEncoding.EncodeToString(signature))
}

type VerifyBody struct {
    Data string `json:"data"`
    Signature []byte `json:"signature"`
}

type VerifyOkResponse struct {
    Correct bool `json:"correct"`
    Message string `json:"message"`
}

type VerifyErrorResponse struct {
    Error string `json:"error"`
}

// @Summary Verify text + signature
// @Description Verify that given signature matches given text, using the public key
// @Accept json
// @Produce json
// @Param body body signer.VerifyBody true "Data-signature pair to verify; signature can be passed as base-64 string"
// @Success 200 {object} signer.VerifyOkResponse
// @Failure 400 {object} signer.VerifyErrorResponse
// @Router /verify [post]
func (d *deps) verify(c *gin.Context) {
    var body VerifyBody
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, VerifyErrorResponse{err.Error()})
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
