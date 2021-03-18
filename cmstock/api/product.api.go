package api

import (
	"cmstock/db"
	"cmstock/interceptor"
	"cmstock/model"
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
	// c.JSON(http.StatusOK, gin.H{"result": []string{"iPhone", "Android", "iPad"}})

	var products []model.Product
    if keyword := c.Query("keyword"); keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name like ?", keyword).Order("created_at DESC").Find(&products)
	}

	
	db.GetDB().Find(&products)
	c.JSON(http.StatusOK, products)
}
