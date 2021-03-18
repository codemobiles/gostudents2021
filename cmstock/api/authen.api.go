package api

import (
	"cmstock/db"
	"cmstock/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(router *gin.Engine) {
	fmt.Println("SetupAuthenAPI is ready")
	authenAPI := router.Group("/api/v2")
	authenAPI.POST("/login", login)
	authenAPI.POST("/register", register)
}

func login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "json binding error"})
		return
	}
	c.JSON(http.StatusOK, user)

}

func register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "json binding error"})
		return
	}

	if err := db.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "insert error"})
		return
	}

	c.JSON(http.StatusOK, user)
}
