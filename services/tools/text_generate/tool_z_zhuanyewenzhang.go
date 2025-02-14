package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolZhuanYeWenZhang struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolZhuanYeWenZhang() *toolZhuanYeWenZhang {
	return &toolZhuanYeWenZhang{
		Message: "写一篇专业类型的文章，主题: {title}, 行业领域: {industry}, 写作风格: {style}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "主题", Placeholder: "文章主题",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "industry", Title: "行业领域", Placeholder: "文章涉及到的领域，服装生产/科技/设备制造等...",
				Value:    "",
				Required: true, ErrorMessage: "请输入行业领域",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value: "正式严谨",
				Options: []schemas.ToolFormItemSelect{
					{Value: "正式严谨", Text: "正式严谨"},
					{Value: "深刻分析", Text: "深刻分析"},
					{Value: "辞藻华丽", Text: "辞藻华丽"},
					{Value: "引经据典", Text: "引经据典"},
					{Value: "论证说理", Text: "论证说理"},
					{Value: "创新拓展", Text: "创新拓展"},
					{Value: "国际视野", Text: "国际视野"},
					{Value: "逻辑严密", Text: "逻辑严密"},
					{Value: "简洁质朴", Text: "简洁质朴"},
				},
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "内容要求",
				Placeholder: "其他要求描述，越详细越好",
				Value:       "",
				Required:    false, ErrorMessage: "请输入内容要求",
			},
		},
	}
}

func (t *toolZhuanYeWenZhang) GetMessage() (result string) {
	return t.Message
}

func (t *toolZhuanYeWenZhang) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolZhuanYeWenZhang) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
