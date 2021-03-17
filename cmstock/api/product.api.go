package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	fmt.Println("SetupProductAPI is ready")
	productAPI := router.Group("/api/v2")
	productAPI.GET("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "product"})
	})
}
