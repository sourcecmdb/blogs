package Handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveUser(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}


// 通过 query 方法进行获取参数
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}