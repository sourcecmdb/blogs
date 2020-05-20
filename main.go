package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app:= gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
