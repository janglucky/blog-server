package web

import "github.com/janglucky/blog-server/common/model"

type Request struct {
	Token  string     `json:"token"`
	User   model.User `json:"user"`
	Params map[string]interface{}
}
