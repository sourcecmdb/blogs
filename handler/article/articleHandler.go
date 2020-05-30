package article

import (
	"github.com/gin-gonic/gin"
	"github.com/sourcecmdb/blogs/model"
	"net/http"
)

func Insert(context *gin.Context){
	article := model.Article{}
	var id  = -1
	if e := context.ShouldBindJSON(&article); e == nil{
	id = article.Insert()
	}
	context.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}
