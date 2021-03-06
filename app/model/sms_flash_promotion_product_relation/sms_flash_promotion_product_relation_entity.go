// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package sms_flash_promotion_product_relation

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table sms_flash_promotion_product_relation.
type Entity struct {
    Id                      int64   `orm:"id,primary"                 json:"id"`                         // 编号          
    FlashPromotionId        int64   `orm:"flash_promotion_id"         json:"flash_promotion_id"`         //               
    FlashPromotionSessionId int64   `orm:"flash_promotion_session_id" json:"flash_promotion_session_id"` // 编号          
    ProductId               int64   `orm:"product_id"                 json:"product_id"`                 //               
    FlashPromotionPrice     float64 `orm:"flash_promotion_price"      json:"flash_promotion_price"`      // 限时购价格    
    FlashPromotionCount     int     `orm:"flash_promotion_count"      json:"flash_promotion_count"`      // 限时购数量    
    FlashPromotionLimit     int     `orm:"flash_promotion_limit"      json:"flash_promotion_limit"`      // 每人限购数量  
    Sort                    int     `orm:"sort"                       json:"sort"`                       // 排序          
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