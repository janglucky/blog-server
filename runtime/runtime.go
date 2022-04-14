package runtime

import (
	"github.com/janglucky/blog-server/initial"
	"github.com/janglucky/blog-server/router"
)
func Execute()  {
	router.StartHttpServer(initial.Settings["port"])
}
