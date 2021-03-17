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
		c.JSON(http.StatusOK, gin.H{"username": "admin", "password": "1234"})
	})

	r.Run(":8081")
}
