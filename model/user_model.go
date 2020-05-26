package model

//type UserModel struct {
//	Email         string `from:"email" binding:"email"`
//	Password      string `from:"password"`
//	PasswordAgain string `from:"password-again" binding:"eqfield=Password"`
//}

type UserModel struct {
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}