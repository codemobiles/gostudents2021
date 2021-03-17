package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})

	r.GET("/register/:username/:password", func(c *gin.Context) {
		username, password := c.Param("username"), c.Param("password")
		c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})

	r.Run(":8081")
}
