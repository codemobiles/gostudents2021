package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyIt(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hey"})
}
