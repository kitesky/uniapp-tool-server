// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameUserScoreLog = "user_score_logs"

// UserScoreLog mapped from table <user_score_logs>
type UserScoreLog struct {
	ID          int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	UserID      int64  `gorm:"column:user_id;type:bigint;comment:用户ID" json:"user_id"` // 用户ID 操作类型 increase-增加 decrease-减少
	Type        string `gorm:"column:type;type:varchar(32);comment:操作类型 increase-增加 decrease-减少" json:"type"`
	Code        string `gorm:"column:code;type:varchar(32);comment:积分来源" json:"code"`  // 积分来源代码
	Amount      int64  `gorm:"column:amount;type:bigint;comment:积分数量" json:"amount"`          // 积分数量
	Title       string `gorm:"column:title;type:varchar(32);comment:标题" json:"title"`      // 标题
	Description string `gorm:"column:description;type:varchar(255);comment:备注" json:"description"` // 备注
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName UserScoreLog's table name
func (*UserScoreLog) TableName() string {
	return TableNameUserScoreLog
}
