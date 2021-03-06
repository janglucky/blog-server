package middleware

import (
	"github.com/kataras/iris/v12"
)

// Cors ...防止跨域
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Method() == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, X-Token")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
