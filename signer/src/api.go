package signer

import (
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
        c.JSON(http.StatusOK, gin.H{"key": cryptographer.Public()})
    })

    type SignBody struct {
        Data string `json:"data"`
    }

    g.POST("/sign", func (c *gin.Context) {
        var body SignBody
        if err := c.ShouldBindJSON(&body); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        signature, err := cryptographer.Sign(body.Data)

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"signature": signature})
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
