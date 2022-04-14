package middleware

import (
	_ "github.com/janglucky/blog-server/common/auth"
	"github.com/kataras/iris/v12"
)

func CheckToken(ctx iris.Context)  {

	ctx.Next()
}
