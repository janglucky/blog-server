package initial

import (
	"github.com/janglucky/blog-server/common/config"
)

var Settings = make(map[string]string)

func init() {
	loadSettings()
}

func loadSettings() {
	err := config.LoadConfig(&Settings, "conf/settings.conf")
	if err != nil {
		panic(err.Error())
	}
}
