package api

import (
	"cmstock/db"
	"cmstock/interceptor"
	"cmstock/model"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})

}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)
		db.GetDB().Model(&product).Update("image", fileName)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}



func deleteProduct(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	db.GetDB().Delete(&model.Product{}, id)
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}