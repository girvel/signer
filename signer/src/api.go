package signer

import (
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"time"
	"strings"

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
// @Description Get the public key used for signature verification in PEM format
// @Produce plain
// @Success 200 {string} string "The key"
// @Router /public [get]
func (d *deps) public(c *gin.Context) {
	c.Header("Cache-Control", fmt.Sprintf("max-age=%v", d.Env.PublicCachingSecs))
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
	if c.GetHeader("Content-Type") != "text/plain" {
	    c.String(http.StatusUnsupportedMediaType, "Content-Type must be text/plain")
		return
	}

	body_bytes, err := c.GetRawData()
	if err != nil {
	    c.String(http.StatusBadRequest, "Error reading body")
		return
	}
    
	datedText := fmt.Sprintf(
		"%s\n\nSigned %s by %s with public key %s/public",
		string(body_bytes),
		time.Now().Format("2006-01-02 15:04:05 MST"),
		d.Env.HolderName,
		d.Env.ServiceAddress,
	)
    signature, err := d.Cryptographer.Sign(datedText)

    if err != nil {
		slog.Error("Error when signing", "error", err)
        c.String(http.StatusInternalServerError, err.Error())
        return
    }

	c.String(http.StatusOK, datedText + "\n\n" + base64.StdEncoding.EncodeToString(signature))
}

type VerifyOkResponse struct {}

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
	if c.GetHeader("Content-Type") != "text/plain" {
	    c.String(http.StatusUnsupportedMediaType, "Content-Type must be text/plain")
		return
	}

	body_bytes, err := c.GetRawData()
	if err != nil {
	    c.String(http.StatusBadRequest, "Error reading body")
		return
	}

	body := string(body_bytes)

	signature_split_i := strings.LastIndex(body, "\n\n")
	if signature_split_i == -1 {
	    c.String(http.StatusBadRequest, "Unable to find signature")
		return
	}

	datedText := body[0:signature_split_i]
	signatureBase64 := body[signature_split_i + 2:]
    
	rawSignature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
	    c.String(http.StatusBadRequest, err.Error())
		return
	}

    err = d.Cryptographer.Verify(datedText, rawSignature)

    if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
    }

    c.Status(http.StatusOK)
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
