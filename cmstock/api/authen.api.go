package api

import (
	"cmstock/db"
	"cmstock/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	user.Password, _ = hashPassword(user.Password)

	if err := db.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "insert error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
