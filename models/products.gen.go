// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name        string    `gorm:"column:name;type:varchar(255);comment:产品名称" json:"name"`                       // 产品名称
	Type        string    `gorm:"column:type;type:varchar(64);comment:产品类型" json:"type"`                        // 产品类型
	Price       float64   `gorm:"column:price;type:decimal(10,2);default:0.00;comment:单价" json:"price"`         // 单价
	MarketPrice float64   `gorm:"column:market_price;type:decimal(10,2);default:0.00;comment:价格" json:"market_price"`      // 市场价格
	Content     string    `gorm:"column:content;type:longtext;comment:详情" json:"content"`                       // 详情
	Status      string    `gorm:"column:status;type:char(32);default:pending;comment:状态 Y上架 N下架" json:"status"` // 状态 Y上架 N下架
	Order       int64     `gorm:"column:order;type:bigint;comment:排序" json:"order"`                             // 排序
	Code        string    `gorm:"column:code;type:varchar(64);comment:产品代码" json:"code"`                        // 产品代码
	Title       string    `gorm:"column:title;type:varchar(255);comment:标题" json:"title"`                       // 标题
	Keywords    string    `gorm:"column:keywords;type:varchar(255);comment:关键词" json:"keywords"`                // 关键词
	Description string    `gorm:"column:description;type:varchar(255);comment:描述" json:"description"`           // 描述
	Extra       string    `gorm:"column:extra;type:varchar(500);comment:标题" json:"extra"`                        // 补充字段信息
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间" json:"created_at"`              // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}
