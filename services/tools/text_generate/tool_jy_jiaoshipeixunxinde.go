package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolJiaoShiPeiXunXinDe struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolJiaoShiPeiXunXinDe() *toolJiaoShiPeiXunXinDe {
	return &toolJiaoShiPeiXunXinDe{
		Message: "写一篇教师培训心得,培训主题:{title},字数限制:{size},其他信息:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "textarea", ValueType: "string", Name: "title", Title: "培训主题", Placeholder: "培训主题",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他信息",
				Placeholder: "补充培训活动的细节信息，例如：教学方法、教学目的、本次活动的感受...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolJiaoShiPeiXunXinDe) GetMessage() (result string) {
	return t.Message
}

func (t *toolJiaoShiPeiXunXinDe) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolJiaoShiPeiXunXinDe) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
