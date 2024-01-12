package middlewares

import (
	"resq-be/utils"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil kode error dari context
		code, exists := c.Get("code")
		if !exists {
			// Jika tidak ada kode, lanjutkan ke middleware berikutnya
			c.Next()
			return
		}

		// Ambil error dari context
		err := c.Errors.Last()
		if err != nil {
			// Jika ada error, ambil pesan error
			message := err.Error()
			// Tampilkan response fail
			utils.Fail(c, code.(int), message)
			c.Abort()
		}
		c.Next()
	}
}
