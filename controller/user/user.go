package user

import (
	"github.com/janglucky/blog-server/common/auth"
	"github.com/janglucky/blog-server/common/model"
	"github.com/janglucky/blog-server/common/web"
	"github.com/kataras/iris/v12"
)

func Verify(ctx iris.Context) {

	req := ctx.Values().Get("reqParams")
	reqParams, ok := req.(web.Request)

	var user model.User

	if err := ctx.ReadJSON(&user); err != nil {
		web.RenderResponse(ctx, web.STATUS_ERROR)
		return
	}

	if err := user.Validate(); err != nil {
		web.RenderResponse(ctx, web.STATUS_VERIFY_ERROR)
		return
	}

	userClaims := &auth.UserClaims{
		Id: user.Id,
		Username: user.Username,
		Password: user.Password,
		Role: user.Role,
	}

	token, err := auth.GenerateToken(userClaims)
	if err != nil {
		web.RenderResponse(ctx, web.STATUS_INTERNAL_ERROR)
		return
	}

	web.RenderResponse(ctx, web.STATUS_OK, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
	return
}

func UserInfo(ctx iris.Context)  {

	web.RenderResponse(ctx, web.STATUS_OK)
	return
}
