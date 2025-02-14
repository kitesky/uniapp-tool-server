package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolYouXiuJiaoShiDaiBiaoFaYan struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolYouXiuJiaoShiDaiBiaoFaYan() *toolYouXiuJiaoShiDaiBiaoFaYan {
	return &toolYouXiuJiaoShiDaiBiaoFaYan{
		Message: "写一篇优秀教师代表发言,发言主题:{title},授课内容{content},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "发言主题", Placeholder: "输入发言主题、关键词",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他要求",
				Placeholder: "输入对发言的其他要求，例如：加入对工作的总结及未来计划...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolYouXiuJiaoShiDaiBiaoFaYan) GetMessage() (result string) {
	return t.Message
}

func (t *toolYouXiuJiaoShiDaiBiaoFaYan) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolYouXiuJiaoShiDaiBiaoFaYan) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
