package auth

import (
	"gf-mall/app/utils/response"

	"github.com/gogf/gf/net/ghttp"
)

//Auth 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	login := false
	if login {
		r.Middleware.Next()
	} else {
		response.JSONExit(r, response.NotLogin, "未登录")
	}
}
