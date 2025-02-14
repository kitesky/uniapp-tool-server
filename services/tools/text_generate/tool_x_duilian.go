package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolDuiLian struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolDuiLian() *toolDuiLian {
	return &toolDuiLian{
		Message: "写一份对联;对联主题:{title};字数限制:{size};上联关键词:{left};下联关键词:{right};是否需要横批:{top}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "对联主题", Placeholder: "主题",
				Value:    "",
				Required: true, ErrorMessage: "请输入对联主题",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "5-50字",
				Value:    8,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "input", ValueType: "string", Name: "left", Title: "上联", Placeholder: "建议2-10个字",
				Value:    "",
				Required: false, ErrorMessage: "请输入上联",
			},
			{
				FieldType: "input", ValueType: "string", Name: "rigth", Title: "下联", Placeholder: "建议2-10个字",
				Value:    "",
				Required: false, ErrorMessage: "请输入下联",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "top", Title: "横批", Placeholder: "",
				Value: "需要",
				Options: []schemas.ToolFormItemSelect{
					{Value: "需要", Text: "需要"},
					{Value: "不需要", Text: "不需要"},
				},
				Required: false, ErrorMessage: "请选择是否需要横批",
			},
		},
	}
}

func (t *toolDuiLian) GetMessage() (result string) {
	return t.Message
}

func (t *toolDuiLian) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolDuiLian) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
