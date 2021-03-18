package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyIt(c *gin.Context) {
	if token := c.Query("token"); token == "1234" {
		c.Next()
	} else {
		c.String(http.StatusOK, "Error no token")
		c.Abort()
	}
}
