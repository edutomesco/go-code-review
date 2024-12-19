package gin

import (
	"github.com/edutomesco/coupons/internal/models/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return errorMiddleware(gin.ErrorTypeAny)
}

func errorMiddleware(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError errors.HttpAppError

			switch err := err.(type) {
			case errors.HttpAppError:
				parsedError = err
			default:
				parsedError = errors.NewHttpAppError(http.StatusInternalServerError, "Internal Server Error")
			}

			c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}

	}
}
