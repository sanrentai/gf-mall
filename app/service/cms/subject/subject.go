package subject

import (
	subjectModel "gf-mall/app/model/cms_subject"
	"gf-mall/app/utils/page"
)

// ListAll 查询所有专题
func ListAll() ([]*subjectModel.Entity, error) {
	return subjectModel.FindAll()
}

// List 分页查询专题
func List(data *subjectModel.SelectPageReq) (*page.Paging, error) {
	return subjectModel.SelectListByPage(data)
}
