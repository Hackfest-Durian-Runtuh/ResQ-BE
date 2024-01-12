package middlewares

import (
	"resq-be/utils"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil kode error dari context
		code, _ := c.Get("code")
		// Ambil error dari context
		err := c.Errors.Last()

		// Jika ada error, ambil pesan error
		message := err.Error()
		// Jika error karena validasi, ambil pesan error dari validasi
		if err.Type == gin.ErrorTypeBind {
			message = err.Err.Error()
		}
		// Tampilkan response fail
		utils.Fail(c, code.(int), message)
		c.Abort()
	}
}
