package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolChanPinJieShao struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolChanPinJieShao() *toolChanPinJieShao {
	return &toolChanPinJieShao{
		Message: "写一篇产品介绍,产品名称:{title},写作风格:{style},字数限制:{size},写作要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "产品名称", Placeholder: "输入产品名称",
				Required: true, ErrorMessage: "请输入产品名称",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "产品信息",
				Placeholder: "输入产品相关信息。例如：产品功能与特点、生产工艺、消费群体...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入产品信息",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value:    "正式严谨",
				Options:  schemas.ToolTextGenerateStyleV4,
				Required: true, ErrorMessage: "请选择风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
		},
	}
}

func (t *toolChanPinJieShao) GetMessage() (result string) {
	return t.Message
}

func (t *toolChanPinJieShao) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolChanPinJieShao) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
