package signer

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAPI(key *rsa.PrivateKey) *gin.Engine {
    g := gin.Default()

    der, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
    if err != nil {
        panic(err.Error())
    }

    pem := pem.EncodeToMemory(&pem.Block{
        Type: "RSA PUBLIC KEY",
        Bytes: der,
    })
    
    g.GET("/public", func (c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"key": string(pem)})
    })

    return g
}
