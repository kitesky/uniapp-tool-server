package schemas

import "app-api/models"

type OrderPageReq struct {
	UserID      int64  `json:"-" binding:"required"`
	OrderStatus string `json:"order_status" form:"order_status"`
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"page_size" form:"page_size"`
	SortField   string `json:"sort_field" form:"sort_field"`
	SortType    string `json:"sort_type" form:"sort_type"`
}

type OrderCreateReq struct {
	UserID       int64   `json:"-" binding:"required"`
	OrderType    string  `json:"order_type" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductCode  string  `json:"product_code"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int32   `json:"quantity"`
	PayType      string  `json:"pay_type"`
}

type OrderRes struct {
	OrderNo       string   `json:"order_no"`
	OrderType     string   `json:"order_type"`
	OrderTypeText string   `json:"order_type_text"`
	PayType       string   `json:"pay_type"`
	PayTypeText   string   `json:"pay_type_text"`
	PayTime       Datetime `json:"pay_time"`
	ProductName   string   `json:"product_name"`
	ProductPrice  float64  `json:"product_price"`
	Quantity      int32    `json:"quantity"`
	Amount        float64  `json:"amount"`
	PayAmount     float64  `json:"pay_amount"`
	PayQrcode     string   `json:"pay_qrcode"`
	Status        string   `json:"status"`
	StatusText    string   `json:"status_text"`
	StatusColor   string   `json:"status_color"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	CreatedAt     Datetime `json:"created_at"`
	UpdatedAt     Datetime `json:"updated_at"`
	ExpiredAt     Datetime `json:"expired_at"`
}

type OrderPageRes struct {
	Total     int             `json:"total"`
	TotalPage int             `json:"total_page"`
	PageSize  int             `json:"page_size"`
	Items     []*models.Order `json:"items"`
}

type OrderNewPageRes struct {
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page"`
	PageSize  int         `json:"page_size"`
	Items     []*OrderRes `json:"items"`
}

type OrderTypeMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// 订单类型
var OrderTypeOptions = map[string]OrderTypeMeta{
	"recharge":    {Title: "充值订单", Description: "充值订单"},
	"vip":         {Title: "VIP订单", Description: "VIP订单"},
	"point":       {Title: "点数充值", Description: "点数充值"},
	"product":     {Title: "消费订单", Description: "消费订单"},
	"application": {Title: "APP应用", Description: "APP应用"},
}

// 订单状态
var OrderStatusOptions = map[string]string{
	"pending":   "待支付",
	"paid":      "已付款",
	"completed": "交易成功",
	"closed":    "订单关闭",
	"canceled":  "订单取消",
	"refunded":  "已退款",
}

// 订单状态
var OrderStatusColorOptions = map[string]string{
	"pending":   "#FF851B",
	"paid":      "#FF4136",
	"completed": "#2ECC40",
	"closed":    "#AAAAAA",
	"canceled":  "#DDDDDD",
	"refunded":  "#39CCCC",
}

// 支付类型
var OrderPayTypeOptions = map[string]string{
	"alipay":  "支付宝",
	"wechat":  "微信支付",
	"balance": "余额支付",
}
