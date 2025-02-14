package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolZhuTiYanJiang struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolZhuTiYanJiang() *toolZhuTiYanJiang {
	return &toolZhuTiYanJiang{
		Message: "写一篇主题演讲,演讲主题:{title},内容风格{style},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "演讲主题", Placeholder: "演讲主题内容",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value:    "正式严谨",
				Options:  schemas.ToolTextGenerateStyleV4,
				Required: true, ErrorMessage: "请选择风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "和主题相关的其他信息内容，例如：要求引用典故、名人名言...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolZhuTiYanJiang) GetMessage() (result string) {
	return t.Message
}

func (t *toolZhuTiYanJiang) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolZhuTiYanJiang) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
