package router

import (
	"gf-mall/app/api/hello"
	_ "gf-mall/app/cms"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})
}
