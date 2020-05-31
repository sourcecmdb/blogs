package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/sourcecmdb/blogs/handler"
	"github.com/sourcecmdb/blogs/handler/article"
	"github.com/sourcecmdb/blogs/middleware"
	"github.com/sourcecmdb/blogs/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path/filepath"
)

func SetupRouter() *gin.Engine {
	//router := gin.Default()
	//router := gin.Default()
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())
	//router.Use(middleware.Logger(),gin.Recovery())
	router.Use(gin.Logger())
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/statics", "./statics/")
	router.StaticFS("/avatar", http.Dir(filepath.Join(utils.RootPath(), "avatar")))
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}
	articleRouter := router.Group("")
	{
		// 通过获取单篇文章
		articleRouter.GET("/article/:id", article.GetOne)
		// 获取所有文章
		articleRouter.GET("/articles", article.GetAll)
		// 添加一篇文章
		articleRouter.POST("/article", article.Insert)
		articleRouter.DELETE("/article/:id", article.DeleteOne)
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
