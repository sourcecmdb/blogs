package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("\n %s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}
