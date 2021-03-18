package api

import (
	"cmstock/db"
	"cmstock/interceptor"
	"cmstock/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	fmt.Println("SetupProductAPI is ready")
	productAPI := router.Group("/api/v2")
	// productAPI.GET("/product", interceptor.VerifyIt, getProduct)
	productAPI.GET("/product", interceptor.JwtVerify, getProduct)
	productAPI.GET("/product/:id", getProductById)
	productAPI.POST("/product", createProduct)
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

func createProduct(c *gin.Context) {
	product := model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreatedAt = time.Now()
	db.GetDB().Create(&product)

	c.JSON(http.StatusOK, product)

}
