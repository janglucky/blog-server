package router

import (
	"github.com/janglucky/blog-server/controller/blog"
	"github.com/janglucky/blog-server/controller/user"
	"github.com/janglucky/blog-server/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"net"
)

func StartHttpServer(port string) {
	app := iris.New()
	app.Use(middleware.Log)
	app.Use(middleware.Cors)

	// 拦截所有options方法
	options := app.Party("/")
	{
		options.Options("*", func(ctx iris.Context) {
			ctx.Next()
		})
	}

	// 博客前台
	app.PartyFunc("/blog", func(party iris.Party) {
		party.Get("/getAllArticles", blog.GetAllArticles)
	})
	
	app.PartyFunc("/user", func(party router.Party) {
		party.Post("/verify", user.Verify)
	})

	// 博客后台
	app.PartyFunc("/admin", func(party iris.Party) {
		party.Use(middleware.CheckToken)
	})

	l, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		panic(err)
	}
	app.Run(iris.Listener(l))
}
