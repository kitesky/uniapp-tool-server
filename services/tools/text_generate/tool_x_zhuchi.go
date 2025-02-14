package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolZhuChi struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolZhuChi() *toolZhuChi {
	return &toolZhuChi{
		Message: "写一篇主持稿，主持题目: {title}, 主持风格: {style}, 字数限制: {size}, 活动信息: {info}, 主持要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "活动事件", Placeholder: "活动事件名称",
				Value:    "",
				Required: true, ErrorMessage: "请输入活动事件名称",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "主持风格", Placeholder: "",
				Value:    "正式严肃",
				Options:  schemas.ToolTextGenerateStyleV3,
				Required: true, ErrorMessage: "请选择主持风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "info", Title: "活动信息", Placeholder: "例如：活动时间、地点、活动内容、活动形式、活动对象等...",
				Value:    "",
				Required: false, ErrorMessage: "请输入活动信息",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "主持要求",
				Placeholder: "提供具体细节内容，让主持稿件更优质。例如：主持场合，听众群体...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolZhuChi) GetMessage() (result string) {
	return t.Message
}

func (t *toolZhuChi) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolZhuChi) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
	result = &schemas.ToolHandlerResponse{
		ContentType: "text",
		Status:      "fail",
	}

	// 请求大模型-deepseek
	resp, _ := deepseek.NewDeepSeek().SendMessage(&schemas.DeepSeekReq{
		Messages:       []schemas.DeepSeekMessage{{Content: message, Role: "user"}},
		Model:          "deepseek-chat",
		MaxTokens:      2048,
		Temperature:    1.5,
		Stream:         false,
		ResponseFormat: schemas.DeepSeekResponseFormat{Type: "text"},
	})

	// 返回数据
	respBody, _ := json.Marshal(resp)
	if len(resp.Choices) > 0 {
		result = &schemas.ToolHandlerResponse{
			Content:      resp.Choices[0].Message.Content,
			ContentType:  "text",
			ResponseBody: string(respBody),
			Status:       "success",
		}
	}

	return
}
