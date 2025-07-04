// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameUserApplication = "user_application"

// UserApplication mapped from table <user_application>
type UserApplication struct {
	ID           int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	UserID       int64     `gorm:"column:user_id;type:bigint;comment:USER ID" json:"user_id"`                            // USER ID
	Name         string    `gorm:"column:name;type:varchar(255);comment:应用名称" json:"name"`                            // 应用名称
	ClientID     string    `gorm:"column:client_id;type:varchar(64);comment:客户端ID" json:"client_id"`                  // 客户端ID
	ClientSecret string    `gorm:"column:client_secret;type:varchar(255);comment:客户端密钥" json:"client_secret"`         // 客户端密钥
	CallbackURL  string    `gorm:"column:callback_url;type:varchar(255);comment:回调地址" json:"callback_url"`            // 回调地址
	Whitelist    string    `gorm:"column:whitelist;type:varchar(255);comment:白名单IP地址" json:"whitelist"`               // 白名单IP地址
	Status       int32     `gorm:"column:status;type:tinyint unsigned;default:1;comment:状态: 0 禁用 1 启用" json:"status"` // 状态: 0 禁用 1 启用
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName UserApplication's table name
func (*UserApplication) TableName() string {
	return TableNameUserApplication
}
