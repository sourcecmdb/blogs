package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", func(context iris.Context) {
		context.HTML("<h1>Hello World</h1>")
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
