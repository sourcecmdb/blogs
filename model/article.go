package model

import (
	"github.com/sourcecmdb/blogs/initDB"
	"log"
)

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (article Article) Insert()  int {
	result, err := initDB.Db.Exec("insert into `article` (type, content) values (?, ?);", article.Type, article.Content)
	if err != nil {
		log.Panicln("文章添加失败", err.Error())
	}
	i, _ := result.LastInsertId()
	return int(i)
}
