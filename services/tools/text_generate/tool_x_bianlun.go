package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolBianLun struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolBianLun() *toolBianLun {
	return &toolBianLun{
		Message: "写一篇辩论稿，辩论题目: {title}, 论述观点: {topic}, 辩论风格: {style}, 字数限制: {size}, 其他要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "辩论主题", Placeholder: "辩论主题",
				Value:    "",
				Required: true, ErrorMessage: "请输入辩论主题",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "topic", Title: "论述观点", Placeholder: "论述观点",
				Value:    "",
				Required: true, ErrorMessage: "请输入论述观点",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "辩论风格", Placeholder: "",
				Value:    "正式",
				Options:  schemas.ToolTextGenerateStyleV2,
				Required: true, ErrorMessage: "请选择风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他要求",
				Placeholder: "提供具体细节内容，让辩论稿件更优质。",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolBianLun) GetMessage() (result string) {
	return t.Message
}

func (t *toolBianLun) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolBianLun) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
