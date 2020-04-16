package response

import (
	"github.com/gogf/gf/net/ghttp"
)

//ErrorCode 错误代码
type ErrorCode string

const (
	//SUCCESS 成功
	SUCCESS ErrorCode = "00"
	//NotLogin 未登录
	NotLogin ErrorCode = "01"

	//OtherError 其他错误
	OtherError ErrorCode = "99"
)

//Resp 数据返回通用JSON数据结构
type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//Success 成功
func Success(r *ghttp.Request, data ...interface{}) {
	JSONExit(r, SUCCESS, "OK", data...)
}

//Fail 失败
func Fail(r *ghttp.Request, msg string, data ...interface{}) {
	JSONExit(r, OtherError, msg, data...)
}

//JSON  标准返回结果数据结构封装。
func JSON(r *ghttp.Request, code ErrorCode, msg string, data ...interface{}) {
	respData := interface{}(nil)
	if len(data) > 0 {
		respData = data[0]
	}
	r.Response.WriteJson(Resp{
		Code: string(code),
		Msg:  msg,
		Data: respData,
	})
}

//JSONExit 返回JSON数据并退出当前HTTP执行函数。
func JSONExit(r *ghttp.Request, code ErrorCode, msg string, data ...interface{}) {
	JSON(r, code, msg, data...)
	r.Exit()
}
