package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolChuangYeJiHuaShu struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolChuangYeJiHuaShu() *toolChuangYeJiHuaShu {
	return &toolChuangYeJiHuaShu{
		Message: "写一份创业计划书,创业方向:{title},创业项目:{subject},写作风格:{style},字数限制:{size},信息补充:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "创业方向", Placeholder: "从事领域。例如：软件开发、餐饮、工业产品...",
				Required: true, ErrorMessage: "请输入创业方向",
			},
			{
				FieldType: "input", ValueType: "string", Name: "subject", Title: "创业项目", Placeholder: "项目名称。例如：XXX火锅店、XXX咖啡店...",
				Required: true, ErrorMessage: "请输入创业项目",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value:    "正式",
				Options:  schemas.ToolTextGenerateStyleV2,
				Required: true, ErrorMessage: "请选择风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "提供一些其他关键信息或者要求。例如：加入市场情况、未来前景、竞品信息...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入写作要求",
			},
		},
	}
}

func (t *toolChuangYeJiHuaShu) GetMessage() (result string) {
	return t.Message
}

func (t *toolChuangYeJiHuaShu) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolChuangYeJiHuaShu) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
