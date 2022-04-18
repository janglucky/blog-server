package middleware

import (
	"encoding/json"
	_ "github.com/janglucky/blog-server/common/auth"
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
	"io/ioutil"
)

func CheckRequest(ctx iris.Context)  {
	var reqParams web.Request

	// 1.读取raw请求
	rawBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		web.RenderResponse(ctx,web.STATUS_ERROR)
		return
	}

	// 2.进行json解码
	err = json.Unmarshal(rawBody, &reqParams)
	if err != nil {
		web.RenderResponse(ctx, web.STATUS_ERROR)
		return
	}
	// 3.保存参数
	ctx.Values().Set("reqParams", reqParams)
	ctx.Next()
}
