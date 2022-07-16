package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogFormatter(param gin.LogFormatterParams) string {

	// custom log format
	return fmt.Sprintf("%s - [%s] %s %s %s %d %s %s %s",
		param.ClientIP,
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
