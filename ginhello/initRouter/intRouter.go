package initRouter

import (
	"fmt"
	"github.com/sourcecmdb/blogs/ginhello/Handler"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// router.GET("/", func(context *gin.Context) {
	// 	context.String(http.StatusOK, "GIN hello world")
	// })

	index := router.Group("/")
	{
		//// 添加 Get 请求路由
		//index.GET("", retHelloGinAndMethod)
		//// 添加 Post 请求路由
		//index.POST("", retHelloGinAndMethod)
		//// 添加 Put 请求路由
		//index.PUT("", retHelloGinAndMethod)
		//// 添加 Delete 请求路由
		//index.DELETE("", retHelloGinAndMethod)
		//// 添加 Patch 请求路由
		//index.PATCH("", retHelloGinAndMethod)
		//// 添加 Head 请求路由
		//index.HEAD("", retHelloGinAndMethod)
		//// 添加 Options 请求路由
		//index.OPTIONS("", retHelloGinAndMethod)
		// 万能方法
		index.Any("",retHelloGinAndMethod)
	}
	UserRouter := router.Group("/user")
	{
		UserRouter.GET("/:name", Handler.SaveUser)
		UserRouter.GET("", Handler.UserSaveByQuery)
	}
	return router
}

func retHelloGinAndMethod(c *gin.Context) {
	c.String(http.StatusOK,"hello gin"+strings.ToLower(c.Request.Method)+"method")
}