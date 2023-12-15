package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s",
			params.ClientIP,
			params.TimeStamp,
			params.MethodColor(),
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
