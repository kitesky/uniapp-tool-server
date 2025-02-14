package schemas

import (
	"app-api/models"
)

type ActivityExchangeReq struct {
	UserID int64 `json:"-" binding:"required"`
	Amount int64 `json:"amount"`
}

type ActivityPageReq struct {
	UserID    int64  `json:"-" binding:"required"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortType  string `json:"sort_type" form:"sort_type"`
	Status    string `json:"status" form:"status"`
}

type ActivityRes struct {
	UUID            string   `json:"uuid"`
	Code            string   `json:"code"`
	Amount          float64  `json:"amount"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Status          string   `json:"status"`
	StatusText      string   `json:"status_text"`
	StatusColor     string   `json:"status_color"`
	RequestBody     string   `json:"request_body"`
	Content         string   `json:"content"`
	ContentType     string   `json:"content_type"`
	ContentTypeText string   `json:"content_type_text"`
	CreatedAt       Datetime `json:"created_at"`
	UpdatedAt       Datetime `json:"updated_at"`
}

type ActivityPageRes struct {
	Total     int                       `json:"total"`
	TotalPage int                       `json:"total_page"`
	PageSize  int                       `json:"page_size"`
	Items     []*models.UserActivityLog `json:"items"`
}

type ActivityNewPageRes struct {
	Total     int            `json:"total"`
	TotalPage int            `json:"total_page"`
	PageSize  int            `json:"page_size"`
	Items     []*ActivityRes `json:"items"`
}

type ActivityStatusMeta struct {
	Text  string `json:"title"`
	Color string `json:"color"`
}

// 类型
var ActivityContentTypeOptions = map[string]OrderTypeMeta{
	"text":  {Title: "文本", Description: "文本"},
	"image": {Title: "图像", Description: "图像"},
	"audio": {Title: "音频", Description: "音频"},
	"video": {Title: "视频", Description: "视频"},
}

// 状态颜色
var ActivityStatusOptions = map[string]ActivityStatusMeta{
	"pending":  {Text: "生成中", Color: "#FF851B"},
	"success":  {Text: "成功", Color: "#2ECC40"},
	"fail":     {Text: "失败", Color: "#AAAAAA"},
	"refunded": {Text: "生成中", Color: "#39CCCC"},
}
