package router

import (
	"github.com/janglucky/blog-server/controller/article"
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

	app.Use(middleware.CheckRequest)

	// 博客前台
	app.PartyFunc("/blog", func(party iris.Party) {
		party.Post("/getAllArticles", blog.GetAllArticles)
	})
	
	app.PartyFunc("/user", func(party router.Party) {
		party.Post("/verify", user.Verify)
	})

	// 管理员相关
	app.PartyFunc("/admin", func(party iris.Party) {
		party.Use(middleware.CheckToken)
		party.Post("/userInfo", user.UserInfo)
		party.Post("/logout", user.Logout)

	})

	// 后台文章相关接口
	app.PartyFunc("/article", func(party iris.Party) {
		party.Use(middleware.CheckToken)
		party.Post("/upload", article.UploadArticle)
		party.Post("/list", article.ListArticle)
		party.Post("/searchTagByKeyword", article.SearchTagByKeyword)
	})


	l, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		panic(err)
	}
	app.Run(iris.Listener(l))
}
