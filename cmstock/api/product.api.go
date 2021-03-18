package api

import (
	"cmstock/interceptor"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	fmt.Println("SetupProductAPI is ready")
	productAPI := router.Group("/api/v2")
	// productAPI.GET("/product", interceptor.VerifyIt, getProduct)
	productAPI.GET("/product", interceptor.JwtVerify, getProduct)
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": []string{"iPhone", "Android", "iPad"}})
}
