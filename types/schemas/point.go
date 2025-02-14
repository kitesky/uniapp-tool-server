package schemas

import (
	"app-api/models"
	"app-api/types/consts"
)

type PointExchangeReq struct {
	UserID int64 `json:"-" binding:"required"`
	Amount int64 `json:"amount"`
}

type PointPageReq struct {
	UserID    int64  `json:"-" binding:"required"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortType  string `json:"sort_type" form:"sort_type"`
}

type PointRes struct {
	Code        string   `json:"code"`
	Amount      int64    `json:"amount"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	TypeText    string   `json:"type_text"`
	TypeColor   string   `json:"type_color"`
	CreatedAt   Datetime `json:"created_at"`
}

type PointPageRes struct {
	Total     int                    `json:"total"`
	TotalPage int                    `json:"total_page"`
	PageSize  int                    `json:"page_size"`
	Items     []*models.UserPointLog `json:"items"`
}

type PointNewPageRes struct {
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page"`
	PageSize  int         `json:"page_size"`
	Items     []*PointRes `json:"items"`
}

// 操作类型
var PointTypeOptions = map[string]string{
	consts.PointActionInc: "+",
	consts.PointActionDec: "-",
}

// 操作类型颜色
var PointTypeColorOptions = map[string]string{
	consts.PointActionInc: "#FF4136",
	consts.PointActionDec: "#2ECC40",
}
