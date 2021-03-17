package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	fmt.Println("SetupTransactionAPI is ready")
	transactionAPI := router.Group("/api/v2")
	transactionAPI.GET("/transaction", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "transaction"})
	})
}
