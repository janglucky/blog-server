package article

import (
	"fmt"
	"github.com/janglucky/blog-server/common/auth"
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func UploadArticle(ctx iris.Context)  {
	uc := ctx.Values().Get("userClaims")
	userClaims, ok := uc.(*auth.UserClaims)
	if !ok {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}
	req := ctx.Values().Get("reqParams")
	reqParams, ok := req.(web.Request)
	if !ok {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}
	article := reqParams.Article
	tags := reqParams.Tags
	fmt.Println(tags)
	article.SetAuthorId(userClaims.Id)
	article.SetCtime()

	err := article.Create()
	if err != nil {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}
	web.RenderResponse(ctx, web.STATUS_OK)
}

func ListArticle(ctx iris.Context)  {

	web.RenderResponse(ctx, web.STATUS_OK)
}