package signer

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cryptographer interface {
    Sign(data string) ([]byte, error)
    Verify(data string, signature []byte) error
    Public() string
}

func CreateAPI(cryptographer Cryptographer) *gin.Engine {
    g := gin.Default()
    
    g.GET("/public", func (c *gin.Context) {
        c.String(http.StatusOK, cryptographer.Public())
    })

    g.POST("/sign", func (c *gin.Context) {
        if c.GetHeader("Content-Type") != "text/plain" {
            c.String(http.StatusUnsupportedMediaType, "Content-Type must be text/plain")
            return
        }

        body, err := c.GetRawData()
        if err != nil {
            c.String(http.StatusBadRequest, "Error reading body")
            return
        }
        
        signature, err := cryptographer.Sign(string(body))

        if err != nil {
            c.String(http.StatusInternalServerError, err.Error())
            return
        }

        c.String(http.StatusOK, base64.StdEncoding.EncodeToString(signature))
    })

    type VerifyBody struct {
        Data string `json:"data"`
        Signature []byte `json:"signature"`
    }

    g.POST("/verify", func (c *gin.Context) {
        var body VerifyBody
        if err := c.ShouldBindJSON(&body); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        err := cryptographer.Verify(body.Data, body.Signature)

        result := gin.H{"correct": err == nil}
        if err != nil {
            result["error"] = err.Error()
        }
        c.JSON(http.StatusOK, result)
    })

    return g
}
