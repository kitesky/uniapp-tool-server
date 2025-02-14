package schemas

import (
	"app-api/models"
	"app-api/types/consts"
)

type BalancePageReq struct {
	UserID    int64  `json:"-" binding:"required"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortType  string `json:"sort_type" form:"sort_type"`
}

type BalanceRes struct {
	Code        string   `json:"code"`
	Amount      float64  `json:"amount"`
	Balance     float64  `json:"balance"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	TypeText    string   `json:"type_text"`
	TypeColor   string   `json:"type_color"`
	CreatedAt   Datetime `json:"created_at"`
}

type BalancePageRes struct {
	Total     int                      `json:"total"`
	TotalPage int                      `json:"total_page"`
	PageSize  int                      `json:"page_size"`
	Items     []*models.UserBalanceLog `json:"items"`
}

type BalanceNewPageRes struct {
	Total     int           `json:"total"`
	TotalPage int           `json:"total_page"`
	PageSize  int           `json:"page_size"`
	Items     []*BalanceRes `json:"items"`
}

// 积分操作类型
var BalanceTypeOptions = map[string]string{
	consts.BalanceActionInc: "+",
	consts.BalanceActionDec: "-",
}

// 积分操作类型颜色
var BalanceTypeColorOptions = map[string]string{
	consts.BalanceActionInc: "#FF4136",
	consts.BalanceActionDec: "#2ECC40",
}
