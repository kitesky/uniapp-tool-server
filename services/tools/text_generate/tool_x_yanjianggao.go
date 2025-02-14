package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolYanJiangGao struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolYanJiangGao() *toolYanJiangGao {
	return &toolYanJiangGao{
		Message: "写一篇演讲稿，演讲题目: {title}, 演讲角色: {role}, 演讲风格: {style}, 字数限制: {size}, 演讲要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "演讲主题", Placeholder: "演讲主题",
				Value:    "",
				Required: true, ErrorMessage: "演讲主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "role", Title: "您的角色", Placeholder: "您的角色",
				Value:    "",
				Required: true, ErrorMessage: "请输入角色",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "演讲风格", Placeholder: "",
				Value:    "正式严肃",
				Options:  schemas.ToolTextGenerateStyleV3,
				Required: true, ErrorMessage: "请选择演讲风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "演讲要求",
				Placeholder: "提供具体细节内容，让演讲稿件更优质。例如：演讲场合，听众群体...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolYanJiangGao) GetMessage() (result string) {
	return t.Message
}

func (t *toolYanJiangGao) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolYanJiangGao) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
