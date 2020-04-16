package cms

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/subject", func(group *ghttp.RouterGroup) {
		group.GET("/listAll", subject.listAll)
	})
}
