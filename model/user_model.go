package model

import (
	"github.com/sourcecmdb/blogs/initDB"
	"log"
)

//type UserModel struct {
//	Email         string `from:"email" binding:"email"`
//	Password      string `from:"password"`
//	PasswordAgain string `from:"password-again" binding:"eqfield=Password"`
//}
//
//type UserModel struct {
//	Id       int64  `form:"id"`
//	Email    string `form:"email"`
//	Password string `form:"password"`
//}
//
//func (user *UserModel) Save() int64 {
//	resful, err := initDB.Db.Exec("INSERT INTO  ginhello.user (email, password) VALUES (?,?);", user.Email, user.Password)
//	if err != nil {
//		log.Panicln("err-> user insert into error", err.Error())
//	}
//	id, err := resful.LastInsertId()
//	if err != nil {
//		log.Panicln("user insert id error", err.Error())
//	}
//	return id
//}
//
//func (user *UserModel) QueryByEmail() UserModel {
//	u := UserModel{}
//	row := initDB.Db.QueryRow("select * from user where email = ?;", user.Email)
//	err := row.Scan(&u.Id, &u.Email, &u.Password)
//	if err != nil {
//		log.Panicln(err)
//	}
//	return u
//}
type UserModel struct {
	Id       int64  `form:"id"`
	Email    string `form:"email" binding:"email"`
	Password string `form:"password" `
}

func (user *UserModel) Save() int64 {
	result, e := initDB.Db.Exec("insert into ginhello.user (email, password) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where email = ?;", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password)
	if e != nil {
		log.Panicln(e)
	}
	return u
}