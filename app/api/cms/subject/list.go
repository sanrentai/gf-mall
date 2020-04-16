package subject

import (
	subjectModel "gf-mall/app/model/cms_subject"
	subjectService "gf-mall/app/service/cms/subject"
	"gf-mall/app/utils/response"

	"github.com/gogf/gf/net/ghttp"
)

//List 根据专题名称分页获取专题
func List(r *ghttp.Request) {

	var req *subjectModel.SelectPageReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		response.Fail(r, err.Error())
	}
	res, err := subjectService.List(req)
	if err != nil {
		response.Fail(r, err.Error())
	}
	response.Success(r, res)
	// response.Fail(r, "Not support this yet!")
}
