package schemas

import (
	"app-api/models"
	"app-api/types/consts"
)

type ScoreExchangeReq struct {
	UserID int64 `json:"-" binding:"required"`
	Amount int64 `json:"amount"`
}

type ScorePageReq struct {
	UserID    int64  `json:"-" binding:"required"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortType  string `json:"sort_type" form:"sort_type"`
}

type ScoreRes struct {
	Code        string   `json:"code"`
	Amount      int64    `json:"amount"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	TypeText    string   `json:"type_text"`
	TypeColor   string   `json:"type_color"`
	CreatedAt   Datetime `json:"created_at"`
}

type ScorePageRes struct {
	Total     int                    `json:"total"`
	TotalPage int                    `json:"total_page"`
	PageSize  int                    `json:"page_size"`
	Items     []*models.UserScoreLog `json:"items"`
}

type ScoreNewPageRes struct {
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page"`
	PageSize  int         `json:"page_size"`
	Items     []*ScoreRes `json:"items"`
}

// 积分操作类型
var ScoreTypeOptions = map[string]string{
	consts.ScoreActionInc: "+",
	consts.ScoreActionDec: "-",
}

// 积分操作类型颜色
var ScoreTypeColorOptions = map[string]string{
	consts.ScoreActionInc: "#FF4136",
	consts.ScoreActionDec: "#2ECC40",
}
