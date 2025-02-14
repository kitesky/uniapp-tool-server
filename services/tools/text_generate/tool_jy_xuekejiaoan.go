package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolXueKeJiaoAn struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolXueKeJiaoAn() *toolXueKeJiaoAn {
	return &toolXueKeJiaoAn{
		Message: "写一篇学科教案,讲课内容:{title},课程时长{minutes}分钟,字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "textarea", ValueType: "string", Name: "title", Title: "讲课内容", Placeholder: "教授学科与教学内容，例如：小学四年级语文课xxx诗词",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "minutes", Title: "课程时间(分钟数)", Placeholder: "输入数字即可，例如:45",
				Required: true, ErrorMessage: "请输入课程时间",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "教案字数", Placeholder: "100-5000字",
				Value:    2000,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "输入对教案的其他要求，例如：增加xxx的分析，xxx的应用...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolXueKeJiaoAn) GetMessage() (result string) {
	return t.Message
}

func (t *toolXueKeJiaoAn) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolXueKeJiaoAn) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
