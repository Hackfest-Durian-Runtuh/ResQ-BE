package middlewares

import (
	"net/http"
	"resq-be/utils"

	"github.com/gin-gonic/gin"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.Fail(c, http.StatusNotFound, "not found")
	}
}
