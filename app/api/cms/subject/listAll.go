package subject

import (
	subjectService "gf-mall/app/service/cms/subject"
	"gf-mall/app/utils/response"

	"github.com/gogf/gf/net/ghttp"
)

//ListAll 获取全部商品专题
func ListAll(r *ghttp.Request) {
	res, err := subjectService.ListAll()
	if err != nil {
		response.Fail(r, err.Error())
	}
	response.Success(r, res)
	// response.Fail(r, "Not support this yet!")
}
