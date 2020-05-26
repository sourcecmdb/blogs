package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sourcecmdb/blogs/model"
	"log"
	"net/http"
)

//func SaveUser(context *gin.Context) {
//	username := context.Param("name")
//	context.String(http.StatusOK, "用户"+username+"已经保存")
//}
//
//// 通过 query 方法进行获取参数
//func UserSaveByQuery(context *gin.Context) {
//	age := context.DefaultQuery("age", "20")
//	username := context.Query("name")
//	// age = context.Query("age")
//	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
//}

func UserRegister(c *gin.Context) {
	////email := c.PostForm("email")
	////password := c.DefaultPostForm("password","123")
	////passwordAgain:= c.DefaultPostForm("password","123")
	////println("email",email,"password",password,"passwordAgin",passwordAgain)
	//var user model.UserModel
	//if err := c.ShouldBind(&user); err != nil {
	//	log.Println("err ->", err.Error())
	//	//println("err ->",err.Error())
	//	c.String(http.StatusBadRequest, "输入不正确")
	//	//return
	//}
	////}else {
	////	log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
	////	c.Redirect(http.StatusMovedPermanently, "/")
	////}
	//////println("email",user.Email,"pasword",user.Password,"password-again",user.PasswordAgain)
	//id := user.Save()
	//log.Println("id is", id)
	//c.Redirect(http.StatusMovedPermanently, "/")\

	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	passwordAgain := c.PostForm("password-again")
	if passwordAgain != user.Password {
		c.String(http.StatusBadRequest, "密码校验无效，两次密码不一致")
		log.Panicln("密码校验无效，两次密码不一致")
	}
	id := user.Save()
	log.Println("id is ", id)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(c *gin.Context) {
	var user model.UserModel
	if err := c.Bind(&user); err != nil {
		log.Panicln("login 绑定错误", err.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"email": u.Email,
		})
	}

}
