package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolJinShengZongJie struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolJinShengZongJie() *toolJinShengZongJie {
	return &toolJinShengZongJie{
		Message: "写一篇晋升总结，主题: {title}, 职位: {job}, 写作风格: {style}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "主题", Placeholder: "主题关键词",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "job", Title: "职位", Placeholder: "职位描述",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value:    "正式",
				Options:  schemas.ToolTextGenerateStyleV2,
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "工作相关的具体信息，例如：工作中遇到的问题及解决方法...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入写作要求",
			},
		},
	}
}

func (t *toolJinShengZongJie) GetMessage() (result string) {
	return t.Message
}

func (t *toolJinShengZongJie) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolJinShengZongJie) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
