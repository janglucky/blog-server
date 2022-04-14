package web

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

type Ack struct {
	Logid  string      `json:"logid"`
	Status int         `json:"status"`
	Msg    string      `json:"msg""`
	Data   interface{} `json:"data""`
}

type EmptyData struct{}

var emptyData EmptyData

func RenderResponse(ctx iris.Context, err error, result ...interface{}) {
	var data interface{}
	status, ok := err.(Status)

	if len(result) > 0 {
		data = result[0]
	} else {
		data = emptyData
	}
	var ack Ack

	logid := fmt.Sprintf("%v", time.Now().Unix())

	if ok {
		ack = Ack{Status: int(status), Msg: status.String(), Data: data, Logid: logid}
	} else {
		ack = Ack{Status: int(STATUS_ERROR), Msg: status.String(), Data: data, Logid: logid}
	}

	body, _ := json.Marshal(ack)
	ctx.Write(body)
}
