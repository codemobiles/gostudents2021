package main

import (
	"app/api"
	_ "fmt"
	_ "os"
	"time"
	_ "time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()
	router.Static("/images", "./uploaded/images")

	// Set up CORS middleware options
	config := cors.Config{
		Origins:        "*",
		RequestHeaders: "Origin, Authorization, Content-Type",

		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	router.Use(cors.Middleware(config))
	api.Setup(router)
	router.Run(":8081")
}
