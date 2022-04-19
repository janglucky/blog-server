package article

import (
	"github.com/janglucky/blog-server/common/model"
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func SearchTagByKeyword(ctx iris.Context) {
	req := ctx.Values().Get("reqParams")
	reqParams, ok := req.(web.Request)
	if !ok {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}
	keywords, ok := reqParams.Params["keyword"]
	if !ok {
		web.RenderResponse(ctx, web.STATUS_PARAMS_INVALID)
		return
	}

	keystr, ok := keywords.(string)
	if !ok {
		web.RenderResponse(ctx, web.STATUS_PARAMS_INVALID)
		return
	}

	var tag model.Tag
	tags, err := tag.SearchByKeyword(keystr)
	if err != nil {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR, err.Error())
		return
	}

	web.RenderResponse(ctx, web.STATUS_OK, tags)
}
