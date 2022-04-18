package web

import "github.com/janglucky/blog-server/common/model"

type Request struct {
	Token         string                 `json:"token"`
	User          model.User             `json:"user"`
	Article       model.Article          `json:"article"`
	Tag           model.Tag              `json:"tag"`
	Params        map[string]interface{} `json:"params"`
}
