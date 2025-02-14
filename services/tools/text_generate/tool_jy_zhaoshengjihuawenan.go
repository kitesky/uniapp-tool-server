package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolZhaoShengJiHuaWenAn struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolZhaoShengJiHuaWenAn() *toolZhaoShengJiHuaWenAn {
	return &toolZhaoShengJiHuaWenAn{
		Message: "写一篇招生计划书,招生主题:{title},招生信息{content},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "招生主题", Placeholder: "主题相关的信息、关键词",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "content", Title: "招生信息",
				Placeholder: "输入机构信息、招生人数、截至日期、专业特色...",
				Required:    false, ErrorMessage: "请输入招生信息",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他要求",
				Placeholder: "其他要求...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolZhaoShengJiHuaWenAn) GetMessage() (result string) {
	return t.Message
}

func (t *toolZhaoShengJiHuaWenAn) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolZhaoShengJiHuaWenAn) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
