package schemas

import "app-api/models"

type UserTransferPageReq struct {
	UserID    int64  `json:"-" binding:"required"`
	Status    string `json:"status" form:"status"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortType  string `json:"sort_type" form:"sort_type"`
}

type UserTransferRes struct {
	Amount      float64  `json:"amount"`
	Account     string   `json:"account"`
	Name        string   `json:"name"`
	OrderNo     string   `json:"order_no"`
	Status      string   `json:"status"`
	StatusText  string   `json:"status_text"`
	StatusColor string   `json:"status_color"`
	PayType     string   `json:"pay_type"`
	PayTypeText string   `json:"pay_type_text"`
	PayOrderNo  string   `json:"pay_order_no"`
	PayTime     string   `json:"pay_time"`
	CreatedAt   Datetime `json:"created_at"`
	UpdatedAt   Datetime `json:"updated_at"`
}

type UserTransferPageRes struct {
	Total     int                    `json:"total"`
	TotalPage int                    `json:"total_page"`
	PageSize  int                    `json:"page_size"`
	Items     []*models.UserTransfer `json:"items"`
}

type UserTransferNewPageRes struct {
	Total     int                `json:"total"`
	TotalPage int                `json:"total_page"`
	PageSize  int                `json:"page_size"`
	Items     []*UserTransferRes `json:"items"`
}

type UserTransferStatusMeta struct {
	Text  string `json:"title"`
	Color string `json:"color"`
}

// 状态颜色
var UserTransferStatusOptions = map[string]UserTransferStatusMeta{
	"pending":  {Text: "待转账", Color: "#FF851B"},
	"success":  {Text: "成功", Color: "#2ECC40"},
	"fail":     {Text: "失败", Color: "#AAAAAA"},
	"refunded": {Text: "已退还", Color: "#39CCCC"},
}

// 支付类型
var UserTransferPayTypeOptions = map[string]string{
	"alipay": "支付宝",
	"wechat": "微信",
}
