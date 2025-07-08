package signer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAPI() *gin.Engine {
    g := gin.Default()
    
    g.GET("/public", func (c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"key": "<NOT A KEY>"})
    })

    return g
}
