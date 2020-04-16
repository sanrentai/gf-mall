package cms

import (
	"gf-mall/app/api/cms/subject"
	"gf-mall/app/middleware/auth"
	"gf-mall/app/middleware/router"
)

func init() {
	g := router.New("admin", "/subject", auth.Auth)
	g.GET("/listAll", "cms:subject:view", subject.ListAll)

}
