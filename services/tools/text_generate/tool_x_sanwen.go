package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolSanWen struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolSanWen() *toolSanWen {
	return &toolSanWen{
		Message: "写一篇散文，写作主题: {title}, 写作风格: {style}, 写作场景: {scene}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "主题、话题、关键词",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "scene", Title: "写作场景", Placeholder: "校园、湖泊、异国他乡、小酒馆...",
				Value:    "",
				Required: true, ErrorMessage: "请输入写作场景",
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
					{Value: "质朴平实", Text: "质朴平实"},
					{Value: "励志奋进", Text: "励志奋进"},
					{Value: "伤感寂寞", Text: "伤感寂寞"},
					{Value: "淡雅清新", Text: "淡雅清新"},
					{Value: "辞藻华丽", Text: "辞藻华丽"},
					{Value: "哲理反思", Text: "哲理反思"},
					{Value: "温馨可爱", Text: "温馨可爱"},
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
				Placeholder: "对于XXX事物的要求，例如：写一篇关于XXX的事...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolSanWen) GetMessage() (result string) {
	return t.Message
}

func (t *toolSanWen) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolSanWen) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
