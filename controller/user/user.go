package user

import (
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func Verify(ctx iris.Context) {
	panic("产生panic错误")
	web.RenderResponse(ctx, web.STATUS_OK)
	return
}
