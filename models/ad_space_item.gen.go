// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameAdSpaceItem = "ad_space_item"

// AdSpaceItem mapped from table <ad_space_item>
type AdSpaceItem struct {
	ID      int64 `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	ItemID  int64 `gorm:"column:item_id;type:bigint;comment:广告内容ID" json:"item_id"`   // 广告内容ID
	SpaceID int64 `gorm:"column:space_id;type:bigint;comment:广告位置ID" json:"space_id"` // 广告位置ID
}

// TableName AdSpaceItem's table name
func (*AdSpaceItem) TableName() string {
	return TableNameAdSpaceItem
}
