package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolGongZuoBaoGao struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolGongZuoBaoGao() *toolGongZuoBaoGao {
	return &toolGongZuoBaoGao{
		Message: "写一篇通知，写作主题: {title}, 写作风格: {style}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "汇报的主题内容，工作任务，完成进度...",
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
				Placeholder: "具体细节，越详细越好。 例如： 工作内容， 完成进度...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入内容要求",
			},
		},
	}
}

func (t *toolGongZuoBaoGao) GetMessage() (result string) {
	return t.Message
}

func (t *toolGongZuoBaoGao) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolGongZuoBaoGao) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
