package main

import (
	"cmstock/api"
	_ "fmt"
	_ "os"
	_ "time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/images", "./uploaded/images")
	api.Setup(router)
	router.Run(":8081")
}
