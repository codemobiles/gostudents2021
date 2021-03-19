package api

import (
	"app/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test() {
	fmt.Println("Test API")
}

func Setup(router *gin.Engine) {
	db.SetupDB()
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
}
