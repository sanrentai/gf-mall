package subject

import "github.com/gogf/gf/net/ghttp"

// 获取全部商品专题
func listAll(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
