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
	productAPI.GET("/product/:id", getProductById)
}

func getProductById(c *gin.Context) {

	var product model.Product
	id := c.Param("id")

	if err := db.GetDB().First(&product, "id=?", id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "not found"})
		return
	} else {
		c.JSON(http.StatusOK, product)
	}

}

func getProduct(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"result": []string{"iPhone", "Android", "iPad"}})

	var products []model.Product
	if keyword := c.Query("keyword"); keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name like ?", keyword).Order("created_at DESC").Find(&products)
	} else {
		db.GetDB().Find(&products)
	}

	c.JSON(http.StatusOK, products)
}
