package main

import (
	_"github.com/sourcecmdb/blogs/docs"
	"github.com/sourcecmdb/blogs/initRouter"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name youngxhu
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
