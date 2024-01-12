package utils

import "github.com/gin-gonic/gin"

func SetError(ctx *gin.Context, err error, message string) {
	ctx.Set("code", ToCode(message))
	ctx.Error(err)
	ctx.Next()
}

const (
	// Error
	ErrInternalServerError = "ERR_INTERNAL_SERVER_ERROR"
	ErrNotFound            = "ERR_NOT_FOUND"
	ErrConflict            = "ERR_CONFLICT"
	ErrBadRequest          = "ERR_BAD_REQUEST"
	ErrForbidden           = "ERR_FORBIDDEN"
	ErrUnauthorized        = "ERR_UNAUTHORIZED"
	ErrUnprocessableEntity = "ERR_UNPROCESSABLE_ENTITY"
	ErrTooManyRequests     = "ERR_TOO_MANY_REQUESTS"
	ErrBadGateway          = "ERR_BAD_GATEWAY"
	ErrServiceUnavailable  = "ERR_SERVICE_UNAVAILABLE"
	ErrGatewayTimeout      = "ERR_GATEWAY_TIMEOUT"
	ErrUnknown             = "ERR_UNKNOWN"
)

func ToCode(errmMessage string) int {
	switch errmMessage {
	case ErrInternalServerError:
		return 500
	case ErrNotFound:
		return 404
	case ErrConflict:
		return 409
	case ErrBadRequest:
		return 400
	case ErrForbidden:
		return 403
	case ErrUnauthorized:
		return 401
	case ErrUnprocessableEntity:
		return 422
	case ErrTooManyRequests:
		return 429
	case ErrBadGateway:
		return 502
	case ErrServiceUnavailable:
		return 503
	case ErrGatewayTimeout:
		return 504
	default:
		return 500
	}
}
