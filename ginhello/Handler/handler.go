package Handler

import (
	"github.com/sourcecmdb/blogs/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveUser(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 通过 query 方法进行获取参数
func UserSaveByQuery(context *gin.Context) {
	age := context.DefaultQuery("age", "20")
	username := context.Query("name")
	// age = context.Query("age")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(c *gin.Context) {
	//email := c.PostForm("email")
	//password := c.DefaultPostForm("password","123")
	//passwordAgain:= c.DefaultPostForm("password","123")
	//println("email",email,"password",password,"passwordAgin",passwordAgain)
	var user model.Usermodel
	if err := c.ShouldBind(&user);err!=nil {
		log.Println("err ->",err.Error())
		//println("err ->",err.Error())
		c.String(http.StatusBadRequest,"输入不正确")
		//return
	}else {
		log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	//println("email",user.Email,"pasword",user.Password,"password-again",user.PasswordAgain)
}
