package middleware

import (
	"github.com/gin-gonic/gin"
)

func GinErrorInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Status() >= 400 {
			// TODO: Insert metrics handler
		}
	}
}
