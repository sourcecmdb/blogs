package model

type Usermodel struct {
	Email         string `from:"email" binding:"email"`
	Password      string `from:"password"`
	PasswordAgain string `from:"password-again" binding:"eqfield=Password"`
}
