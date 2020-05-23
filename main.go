package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sourcecmdb/blogs/ginhello/initRouter"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	gin.SetMode(gin.DebugMode)  // debug模式
	gin.SetMode(gin.TestMode) // test测试模式
 	app := initRouter.SetupRouter()
	_ = app.Run(":8080",)
}
