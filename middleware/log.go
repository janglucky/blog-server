package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

// Log ...防止跨域
func Log(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ctx.Next()
}
