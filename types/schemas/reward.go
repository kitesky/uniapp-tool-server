package schemas

import (
	"app-api/models"
	"app-api/types/consts"
)

type RewardExchangeReq struct {
	UserID    int64   `json:"-" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	CashType  string  `json:"cash_type" binding:"required"`
	PayType   string  `json:"pay_type"`
	PaymentID int64   `json:"payment_id"`
	ProductID int64   `json:"product_id"`
}

type RewardPageReq struct {
	UserID    int64  `json:"-" binding:"required"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortType  string `json:"sort_type" form:"sort_type"`
}

type RewardRes struct {
	Code        string   `json:"code"`
	Amount      float64  `json:"amount"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	TypeText    string   `json:"type_text"`
	TypeColor   string   `json:"type_color"`
	CreatedAt   Datetime `json:"created_at"`
}

type RewardPageRes struct {
	Total     int                     `json:"total"`
	TotalPage int                     `json:"total_page"`
	PageSize  int                     `json:"page_size"`
	Items     []*models.UserRewardLog `json:"items"`
}

type RewardNewPageRes struct {
	Total     int          `json:"total"`
	TotalPage int          `json:"total_page"`
	PageSize  int          `json:"page_size"`
	Items     []*RewardRes `json:"items"`
}

// 积分操作类型
var RewardTypeOptions = map[string]string{
	consts.RewardActionInc: "+",
	consts.RewardActionDec: "-",
}

// 积分操作类型颜色
var RewardTypeColorOptions = map[string]string{
	consts.RewardActionInc: "#FF4136",
	consts.RewardActionDec: "#2ECC40",
}
