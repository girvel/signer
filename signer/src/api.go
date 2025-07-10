package signer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
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

    pssOptions := rsa.PSSOptions{
        SaltLength: rsa.PSSSaltLengthAuto,
        Hash: crypto.SHA256,
    }

    type SignBody struct {
        Data string `json:"data"`
    }

    g.POST("/sign", func (c *gin.Context) {
        var body SignBody
        if err := c.ShouldBindJSON(&body); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        hashed := sha256.Sum256([]byte(body.Data))

        signature, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, hashed[:], &pssOptions)

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
        }

        hashed := sha256.Sum256([]byte(body.Data))

        err := rsa.VerifyPSS(
            &key.PublicKey,
            crypto.SHA256,
            hashed[:],
            body.Signature,
            &pssOptions,
        )

        result := gin.H{"correct": err == nil}
        if err != nil {
            result["error"] = err.Error()
        }
        c.JSON(http.StatusOK, result)
    })

    return g
}
