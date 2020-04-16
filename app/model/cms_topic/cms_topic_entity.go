// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package cms_topic

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table cms_topic.
type Entity struct {
    Id             int64       `orm:"id,primary"      json:"id"`              //           
    CategoryId     int64       `orm:"category_id"     json:"category_id"`     //           
    Name           string      `orm:"name"            json:"name"`            //           
    CreateTime     *gtime.Time `orm:"create_time"     json:"create_time"`     //           
    StartTime      *gtime.Time `orm:"start_time"      json:"start_time"`      //           
    EndTime        *gtime.Time `orm:"end_time"        json:"end_time"`        //           
    AttendCount    int         `orm:"attend_count"    json:"attend_count"`    // 参与人数  
    AttentionCount int         `orm:"attention_count" json:"attention_count"` // 关注人数  
    ReadCount      int         `orm:"read_count"      json:"read_count"`      //           
    AwardName      string      `orm:"award_name"      json:"award_name"`      // 奖品名称  
    AttendType     string      `orm:"attend_type"     json:"attend_type"`     // 参与方式  
    Content        string      `orm:"content"         json:"content"`         // 话题内容  
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