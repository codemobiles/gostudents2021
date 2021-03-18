package api

import (
	"main/db"
	"main/interceptor"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction" /* interceptor.JwtVerify,*/, getTransaction)
		transactionAPI.POST("/transaction", interceptor.JwtVerify, createTransaction)
	}
}

type TransactionResult struct {
	Change        float64 `json:"change"`
	CreatedAt     string  `json:"created_at"`
	ID            uint    `json:"id"`
	OrderList     string  `json:"order_list"`
	Paid          float64 `json:"paid"`
	PaymentDetail string  `json:"payment_detail"`
	PaymentType   string  `json:"payment_type"`
	Total         float64 `json:"total"`
	Staff         string  `json:"staff_id"`
}

func getTransaction(c *gin.Context) {
	// var result []map[string]interface{}
	var result []TransactionResult
	db.GetDB().Debug().Raw("SELECT transactions.id, total, paid, change, payment_type, payment_detail, order_list, users.username as Staff, transactions.created_at FROM transactions join users on transactions.staff_id = users.id order by transactions.created_at DESC", nil).Scan(&result)
	c.JSON(200, result)
}

func createTransaction(c *gin.Context) {
	
}
