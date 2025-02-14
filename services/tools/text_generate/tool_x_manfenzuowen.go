package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolManFenZuoWen struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolManFenZuoWen() *toolManFenZuoWen {
	return &toolManFenZuoWen{
		Message: "写一篇作文，写作题目: {title}, 写作风格: {style}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "作文题目", Placeholder: "作文题目",
				Value:    "",
				Required: true, ErrorMessage: "请输入题目",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "writing_style", Title: "文章体裁", Placeholder: "",
				Value: "记叙文",
				Options: []schemas.ToolFormItemSelect{
					{Value: "记叙文", Text: "记叙文"},
					{Value: "说明文", Text: "说明文"},
					{Value: "议论文", Text: "议论文"},
					{Value: "应用文", Text: "应用文"},
					{Value: "诗歌", Text: "诗歌"},
					{Value: "小说", Text: "小说"},
					{Value: "戏剧", Text: "戏剧"},
					{Value: "散文", Text: "散文"},
				},
				Required: true, ErrorMessage: "请选择文章体裁",
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
				Placeholder: "例如：要求引用古籍、诗词信息、名人名言...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入内容要求",
			},
		},
	}
}

func (t *toolManFenZuoWen) GetMessage() (result string) {
	return t.Message
}

func (t *toolManFenZuoWen) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolManFenZuoWen) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
