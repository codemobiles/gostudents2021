package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(router *gin.Engine) {
	fmt.Println("SetupAuthenAPI is ready")
	authenAPI := router.Group("/api/v2")
	authenAPI.GET("/authen", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "authen"})
	})
}
