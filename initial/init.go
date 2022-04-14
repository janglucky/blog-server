package initial

import (
	"github.com/janglucky/blog-server/common/config"
	"github.com/janglucky/blog-server/common/model"
)

var Settings = make(map[string]string)

func init() {
	loadSettings()
	loadMysql()
}

func loadSettings() {
	err := config.LoadConfig(&Settings, "conf/settings.conf")
	if err != nil {
		panic(err.Error())
	}
}

func loadMysql()  {
	model.NewDb(Settings)
}