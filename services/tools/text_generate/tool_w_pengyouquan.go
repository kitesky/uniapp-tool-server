package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolPengYouQuan struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolPengYouQuan() *toolPengYouQuan {
	return &toolPengYouQuan{
		Message: "写一篇微信朋友圈文案,主题:{title},生成数量:{count},受众群体:{target},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "例如: 日常生活、美食旅行、工作想法...",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "count", Title: "生成数量", Placeholder: "生成数量",
				Value: "1条",
				Options: []schemas.ToolFormItemSelect{
					{Value: "1条", Text: "1条"},
					{Value: "3条", Text: "3条"},
					{Value: "5条", Text: "5条"},
					{Value: "10条", Text: "10条"},
				},
				Required: true, ErrorMessage: "请选择生成条数",
			},
			{
				FieldType: "input", ValueType: "string", Name: "target", Title: "受众群体", Placeholder: "例如: 学生、00后、职业宝妈...",
				Required: true, ErrorMessage: "请输入受众群体",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "例如: 能激发阅读者兴趣、能够引起读者的共鸣...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolPengYouQuan) GetMessage() (result string) {
	return t.Message
}

func (t *toolPengYouQuan) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolPengYouQuan) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
