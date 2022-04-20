package article

import (
	"fmt"
	"github.com/janglucky/blog-server/common/auth"
	"github.com/janglucky/blog-server/common/model"
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

	for i := range tags {
		if tags[i].Id <= 0 {
			tags[i].Id = 0
			if err := tags[i].Create(); err != nil {
				// todo 错误日志入库
				fmt.Println(err.Error())
			}
		}
	}
	err := article.Create()
	if err != nil {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}
	// 文章和标签关系表新增记录
	var articleTagRel model.ArticleTagRel
	for i := range tags {
		articleTagRel.ArticleId = article.Id
		articleTagRel.TagId = tags[i].Id
		articleTagRel.Create()
	}

	web.RenderResponse(ctx, web.STATUS_OK)
}

func ListArticle(ctx iris.Context)  {

	web.RenderResponse(ctx, web.STATUS_OK)
}