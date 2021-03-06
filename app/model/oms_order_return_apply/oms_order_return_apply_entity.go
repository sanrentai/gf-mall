// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package oms_order_return_apply

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table oms_order_return_apply.
type Entity struct {
    Id               int64       `orm:"id,primary"         json:"id"`                 //                                                       
    OrderId          int64       `orm:"order_id"           json:"order_id"`           // 订单id                                                
    CompanyAddressId int64       `orm:"company_address_id" json:"company_address_id"` // 收货地址表id                                          
    ProductId        int64       `orm:"product_id"         json:"product_id"`         // 退货商品id                                            
    OrderSn          string      `orm:"order_sn"           json:"order_sn"`           // 订单编号                                              
    CreateTime       *gtime.Time `orm:"create_time"        json:"create_time"`        // 申请时间                                              
    MemberUsername   string      `orm:"member_username"    json:"member_username"`    // 会员用户名                                            
    ReturnAmount     float64     `orm:"return_amount"      json:"return_amount"`      // 退款金额                                              
    ReturnName       string      `orm:"return_name"        json:"return_name"`        // 退货人姓名                                            
    ReturnPhone      string      `orm:"return_phone"       json:"return_phone"`       // 退货人电话                                            
    Status           int         `orm:"status"             json:"status"`             // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝  
    HandleTime       *gtime.Time `orm:"handle_time"        json:"handle_time"`        // 处理时间                                              
    ProductPic       string      `orm:"product_pic"        json:"product_pic"`        // 商品图片                                              
    ProductName      string      `orm:"product_name"       json:"product_name"`       // 商品名称                                              
    ProductBrand     string      `orm:"product_brand"      json:"product_brand"`      // 商品品牌                                              
    ProductAttr      string      `orm:"product_attr"       json:"product_attr"`       // 商品销售属性：颜色：红色；尺码：xl;                   
    ProductCount     int         `orm:"product_count"      json:"product_count"`      // 退货数量                                              
    ProductPrice     float64     `orm:"product_price"      json:"product_price"`      // 商品单价                                              
    ProductRealPrice float64     `orm:"product_real_price" json:"product_real_price"` // 商品实际支付单价                                      
    Reason           string      `orm:"reason"             json:"reason"`             // 原因                                                  
    Description      string      `orm:"description"        json:"description"`        // 描述                                                  
    ProofPics        string      `orm:"proof_pics"         json:"proof_pics"`         // 凭证图片，以逗号隔开                                  
    HandleNote       string      `orm:"handle_note"        json:"handle_note"`        // 处理备注                                              
    HandleMan        string      `orm:"handle_man"         json:"handle_man"`         // 处理人员                                              
    ReceiveMan       string      `orm:"receive_man"        json:"receive_man"`        // 收货人                                                
    ReceiveTime      *gtime.Time `orm:"receive_time"       json:"receive_time"`       // 收货时间                                              
    ReceiveNote      string      `orm:"receive_note"       json:"receive_note"`       // 收货备注                                              
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}