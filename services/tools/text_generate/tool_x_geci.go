package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolGeCi struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolGeCi() *toolGeCi {
	return &toolGeCi{
		Message: "写一篇歌词，写作主题: {title}, 情景: {scene}, 写作风格: {style}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "主题、话题、关键词",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "scene", Title: "情景", Placeholder: "校园、毕业、分手、旅行、小酒馆...",
				Value:    "",
				Required: true, ErrorMessage: "请输入情景",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value: "简洁明了",
				Options: []schemas.ToolFormItemSelect{
					{Value: "简洁明了", Text: "简洁明了"},
					{Value: "平实质朴", Text: "平实质朴"},
					{Value: "幽默风趣", Text: "幽默风趣"},
					{Value: "浪漫幻想", Text: "浪漫幻想"},
					{Value: "严肃正式", Text: "严肃正式"},
					{Value: "夸张猎奇", Text: "夸张猎奇"},
					{Value: "细腻抒情", Text: "细腻抒情"},
					{Value: "淡雅清丽", Text: "淡雅清丽"},
					{Value: "犀利毒舌", Text: "犀利毒舌"},
					{Value: "气势磅礴", Text: "气势磅礴"},
					{Value: "励志成长", Text: "励志成长"},
					{Value: "伤感怀念", Text: "伤感怀念"},
					{Value: "愉悦欢喜", Text: "愉悦欢喜"},
					{Value: "孤独寂寞", Text: "孤独寂寞"},
					{Value: "温馨温暖", Text: "温馨温暖"},
					{Value: "科技科幻", Text: "科技科幻"},
					{Value: "感恩感谢", Text: "感恩感谢"},
				},
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "例如：押韵流畅、韵律美妙、语调自然、结构清晰、内容丰富...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolGeCi) GetMessage() (result string) {
	return t.Message
}

func (t *toolGeCi) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolGeCi) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
