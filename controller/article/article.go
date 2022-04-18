package article

import (
	"fmt"
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func UploadArticle(ctx iris.Context)  {

	req := ctx.Values().Get("reqParams")
	reqParams, ok := req.(web.Request)
	if !ok {
		fmt.Println(reqParams)

		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}
	article := reqParams.Article
	err := article.Create()
	if err != nil {
		fmt.Println(err.Error())
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}
	web.RenderResponse(ctx, web.STATUS_OK)
}

func ListArticle(ctx iris.Context)  {

	web.RenderResponse(ctx, web.STATUS_OK)
}