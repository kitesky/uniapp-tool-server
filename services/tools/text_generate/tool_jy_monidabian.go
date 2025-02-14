package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolMoNiDaBian struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolMoNiDaBian() *toolMoNiDaBian {
	return &toolMoNiDaBian{
		Message: "写一篇毕业论文模拟答辩,答辩题目:{title},研究内容:{content},研究摘要:{desc},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "答辩题目", Placeholder: "毕业论文答辩题目",
				Required: true, ErrorMessage: "请输入答辩题目",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "content", Title: "研究内容", Placeholder: "研究内容",
				Required: false, ErrorMessage: "请输入研究内容",
			},
			{
				FieldType: "input", ValueType: "string", Name: "desc", Title: "研究摘要", Placeholder: "研究摘要",
				Required: false, ErrorMessage: "请输入研究摘要",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "毕业答辩的其他要求，例如：研究课题的创新点、主要内容和结论，背景意义和现状...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolMoNiDaBian) GetMessage() (result string) {
	return t.Message
}

func (t *toolMoNiDaBian) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolMoNiDaBian) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
