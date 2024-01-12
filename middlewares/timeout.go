package middlewares

import (
	"net/http"
	"resq-be/utils"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout(second int) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(second)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			utils.Fail(c, http.StatusRequestTimeout, "request timeout")
		}),
	)
}
