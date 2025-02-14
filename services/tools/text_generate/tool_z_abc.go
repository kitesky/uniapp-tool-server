package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolABC struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolABC() *toolABC {
	return &toolABC{
		Message: "写一篇年终总结，职业: {title}, 工作描述: {content}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "自定义标题", Placeholder: "请输入标题",
				Required: true, ErrorMessage: "请输入标题",
			},
			{
				FieldType: "number", ValueType: "int", Name: "test", Title: "测试字段", Placeholder: "请输入测试",
				Required: true, ErrorMessage: "请输入测试",
			},
			{
				FieldType: "select", ValueType: "string", Name: "select", Title: "自定义select", Placeholder: "请输入测试",
				Options: []schemas.ToolFormItemSelect{
					{Value: "选项1", Text: "选项1"},
					{Value: "选项2", Text: "选项2"},
				},
				Required: true, ErrorMessage: "请输入测试",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "content", Title: "自定义内容", Placeholder: "请输入内容",
				Required: true, ErrorMessage: "请输入内容",
			},
		},
	}
}

func (t *toolABC) GetMessage() (result string) {
	return t.Message
}

func (t *toolABC) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolABC) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
