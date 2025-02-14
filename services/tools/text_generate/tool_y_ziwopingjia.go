package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolZiWoPingJia struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolZiWoPingJia() *toolZiWoPingJia {
	return &toolZiWoPingJia{
		Message: "写一份自我评价,职业:{role},主题:{title},写作要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "role", Title: "职业", Placeholder: "例如: 教师、程序员、设计师",
				Value:    "",
				Required: true, ErrorMessage: "请输入职业",
			},
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "主题题目", Placeholder: "建议的主题、题目、事情",
				Value:    "",
				Required: true, ErrorMessage: "请输入写作主题",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "自我评价的其他要求。例如: 加入反思、下一步计划...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolZiWoPingJia) GetMessage() (result string) {
	return t.Message
}

func (t *toolZiWoPingJia) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolZiWoPingJia) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
