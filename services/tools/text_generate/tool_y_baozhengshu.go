package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolBaoZhengShu struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolBaoZhengShu() *toolBaoZhengShu {
	return &toolBaoZhengShu{
		Message: "写一份保证书,写作主题:{title},写作风格:{style},字数限制:{size},写作要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "主题、事件、事项",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value:    "正式严谨",
				Options:  schemas.ToolTextGenerateStyleV1,
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "其他要求",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolBaoZhengShu) GetMessage() (result string) {
	return t.Message
}

func (t *toolBaoZhengShu) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolBaoZhengShu) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
