package user

import (
	"fmt"
	"github.com/janglucky/blog-server/common/model"
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func Verify(ctx iris.Context) {

	var user model.User

	if err := ctx.ReadJSON(&user); err != nil {
		web.RenderResponse(ctx, web.STATUS_ERROR, err.Error())
		return
	}
	fmt.Println(user)

	if err := user.Validate(); err != nil {
		fmt.Println(err)
		web.RenderResponse(ctx, web.STATUS_VERIFY_ERROR, err.Error())

	}

	fmt.Println(user)
	//model.User.Validate();
	web.RenderResponse(ctx, web.STATUS_OK)
	return
}
