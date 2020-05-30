package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sourcecmdb/blogs/initRouter"
)

func main() {

	router := initRouter.SetupRouter()

	//gin.SetMode(gin.DebugMode) // debug模式
	gin.SetMode(gin.ReleaseMode)
	_ = router.Run(":8080")
}
