package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolTuoKouXiu struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolTuoKouXiu() *toolTuoKouXiu {
	return &toolTuoKouXiu{
		Message: "写一篇脱口秀演讲稿，写作主题: {title}, 写作风格: {style}, 脱口秀类型: {type}, 字数限制: {size}, 观点: {topic}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "脱口秀主题",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value: "幽默风趣",
				Options: []schemas.ToolFormItemSelect{
					{Value: "幽默风趣", Text: "幽默风趣"},
					{Value: "嘲讽批判", Text: "嘲讽批判"},
					{Value: "犀利毒舌", Text: "犀利毒舌"},
					{Value: "浪漫抒情", Text: "浪漫抒情"},
					{Value: "哲思说理", Text: "哲思说理"},
					{Value: "夸张个性", Text: "夸张个性"},
					{Value: "温馨温暖", Text: "温馨温暖"},
					{Value: "简洁明了", Text: "简洁明了"},
					{Value: "激情昂扬", Text: "激情昂扬"},
				},
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "topic", Title: "主要观点",
				Placeholder: "对XXX某的看法，热门事件的观点...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入主要观点",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "例如：网络流行梗...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolTuoKouXiu) GetMessage() (result string) {
	return t.Message
}

func (t *toolTuoKouXiu) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolTuoKouXiu) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
