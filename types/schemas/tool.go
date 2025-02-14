package schemas

import "app-api/models"

// 表单类型-select
type ToolFormItemSelect struct {
	Text    string `json:"text"`
	Value   string `json:"value"`
	Disable string `json:"disable"`
}

// 表单类型
type ToolFormItem struct {
	FieldType    string               `json:"field_type"` // 字段类型 input,textarea,select,radio,checkbox,file,image
	Title        string               `json:"title"`
	Name         string               `json:"name"`
	Value        interface{}          `json:"value"`
	ValueType    string               `json:"value_type"`
	Placeholder  string               `json:"placeholder"`
	Options      []ToolFormItemSelect `json:"options"`
	Required     bool                 `json:"required"`
	Min          int                  `json:"min"`
	Max          int                  `json:"max"`
	ErrorMessage string               `json:"error_message"`
}

var ToolTextGenerateStyleV1 = []ToolFormItemSelect{
	{Value: "正式严谨", Text: "正式严谨"},
	{Value: "简洁明了", Text: "简洁明了"},
	{Value: "认真严肃", Text: "认真严肃"},
	{Value: "辞藻华丽", Text: "辞藻华丽"},
	{Value: "深刻分析", Text: "深刻分析"},
	{Value: "轻松幽默", Text: "轻松幽默"},
	{Value: "感恩答谢", Text: "感恩答谢"},
	{Value: "哲理感悟", Text: "哲理感悟"},
	{Value: "细腻抒情", Text: "细腻抒情"},
}

var ToolTextGenerateStyleV2 = []ToolFormItemSelect{
	{Value: "正式", Text: "正式"},
	{Value: "非正式", Text: "非正式"},
	{Value: "幽默", Text: "幽默"},
	{Value: "严肃", Text: "严肃"},
	{Value: "简洁", Text: "简洁"},
	{Value: "辞藻华丽", Text: "辞藻华丽"},
	{Value: "深刻", Text: "深刻"},
	{Value: "感恩", Text: "感恩"},
	{Value: "感悟", Text: "感悟"},
	{Value: "抒情", Text: "抒情"},
	{Value: "说理", Text: "说理"},
	{Value: "夸张", Text: "夸张"},
	{Value: "科技感", Text: "科技感"},
	{Value: "说理", Text: "说理"},
}

var ToolTextGenerateStyleV3 = []ToolFormItemSelect{
	{Value: "正式严肃", Text: "正式严肃"},
	{Value: "激情昂扬", Text: "激情昂扬"},
	{Value: "辞藻华丽", Text: "辞藻华丽"},
	{Value: "简洁平实", Text: "简洁平实"},
	{Value: "浪漫抒情", Text: "浪漫抒情"},
	{Value: "幽默诙谐", Text: "幽默诙谐"},
	{Value: "分析说理", Text: "分析说理"},
	{Value: "温馨感恩", Text: "温馨感恩"},
	{Value: "客观论证", Text: "客观论证"},
}

var ToolTextGenerateStyleV4 = []ToolFormItemSelect{
	{Value: "正式严谨", Text: "正式严谨"},
	{Value: "简明实用", Text: "简明实用"},
	{Value: "励志哲思", Text: "励志哲思"},
	{Value: "深刻犀利", Text: "深刻犀利"},
	{Value: "风趣调侃", Text: "风趣调侃"},
	{Value: "自然清新", Text: "自然清新"},
	{Value: "辞藻华丽", Text: "辞藻华丽"},
	{Value: "客观论证", Text: "客观论证"},
	{Value: "浪漫抒情", Text: "浪漫抒情"},
	{Value: "夸张猎奇", Text: "夸张猎奇"},
	{Value: "伤感低落", Text: "伤感低落"},
	{Value: "科技奇幻", Text: "科技奇幻"},
}

// 请求返回结果
type ToolHandlerResponse struct {
	UUID         string `json:"uuid"`
	Status       string `json:"status"`
	ContentType  string `json:"content_type"` // text,image,audio,video
	Content      string `json:"content"`      // 结果内容
	ResponseBody string `json:"response_body"`
}

type ToolHandlerReq struct {
	UserID int64  `json:"-" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Data   string `json:"data" binding:"required"`
}

type ToolPageReq struct {
	KeyWord   string `json:"keyword" form:"keyword"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortType  string `json:"sort_type" form:"sort_type"`
}

type ToolRes struct {
	ID          int64           `json:"id"`
	Code        string          `json:"code"`
	TaxonomyID  int64           `json:"taxonomy_id"`
	Name        string          `json:"name"`
	Icon        string          `json:"icon"`
	Url         string          `json:"url"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Content     string          `json:"content"`
	Price       float64         `json:"price"`
	Recommend   string          `json:"recommend"`
	FormSchemas []*ToolFormItem `json:"form_schemas"`
}

type ToolPageRes struct {
	Total     int               `json:"total"`
	TotalPage int               `json:"total_page"`
	PageSize  int               `json:"page_size"`
	Items     []*models.AppTool `json:"items"`
}

type ToolNewPageRes struct {
	Total     int        `json:"total"`
	TotalPage int        `json:"total_page"`
	PageSize  int        `json:"page_size"`
	Items     []*ToolRes `json:"items"`
}
