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
	api.Test()
	router.Run(":8081")
}
