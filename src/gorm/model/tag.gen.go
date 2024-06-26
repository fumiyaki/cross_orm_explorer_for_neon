// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTag = "Tag"

// Tag mapped from table <Tag>
type Tag struct {
	ID        int32     `gorm:"column:id;primaryKey;default:nextval('"Tag_id_seq"" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null" json:"updatedAt"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
