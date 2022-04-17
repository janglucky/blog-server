package middleware

import (
	"github.com/janglucky/blog-server/common/auth"
	_ "github.com/janglucky/blog-server/common/auth"
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func CheckToken(ctx iris.Context)  {


	req := ctx.Values().Get("reqParams")
	reqParams, ok := req.(web.Request)

	if !ok {
		web.RenderResponse(ctx, web.STATUS_NOT_LOGIN)
		return
	}

	token := reqParams.Token
	userClaims, err := auth.ParseToken(token)
	if err != nil {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR, err)
		return
	}

	ctx.Values().Set("userClaims", userClaims)
	ctx.Next()
}
