package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolGongSiJieShao struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolGongSiJieShao() *toolGongSiJieShao {
	return &toolGongSiJieShao{
		Message: "写一篇公司介绍,公司名称:{title},公司信息:{content},写作风格:{style},字数限制:{size}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "公司名称", Placeholder: "输入公司名称",
				Required: true, ErrorMessage: "请输入公司名称",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "content", Title: "公司信息",
				Placeholder: "输入公司信息。例如：公司主营业务、产品信息、团队人员等、企业文化...",
				Value:       "",
				Required:    true, ErrorMessage: "请输入写作要求",
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

func (t *toolGongSiJieShao) GetMessage() (result string) {
	return t.Message
}

func (t *toolGongSiJieShao) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolGongSiJieShao) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
