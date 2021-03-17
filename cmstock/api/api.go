package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test() {
	fmt.Println("Test API")
}

func Setup(router *gin.Engine) {
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupProductAPI(router)
}
