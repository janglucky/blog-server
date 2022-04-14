package blog

import (
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func GetAllArticles(ctx iris.Context)  {
	web.RenderResponse(ctx, web.STATUS_OK)
	return
}
