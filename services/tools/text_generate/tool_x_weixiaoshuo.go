package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolWeiXiaoShuo struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolWeiXiaoShuo() *toolWeiXiaoShuo {
	return &toolWeiXiaoShuo{
		Message: "写一篇微小说，写作主题: {title}, 写作风格: {style}, 写作背景: {scene}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "主题、题目，话题",
				Value:    "",
				Required: true, ErrorMessage: "请输入写作主题",
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
				},
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "select", ValueType: "string", Name: "scene", Title: "写作背景", Placeholder: "",
				Value: "日常生活",
				Options: []schemas.ToolFormItemSelect{
					{Value: "日常生活", Text: "日常生活"},
					{Value: "都市情感", Text: "都市情感"},
					{Value: "青春校园", Text: "青春校园"},
					{Value: "古风武侠", Text: "古风武侠"},
					{Value: "家庭亲情", Text: "家庭亲情"},
					{Value: "未来科技", Text: "未来科技"},
					{Value: "寓言童话", Text: "寓言童话"},
					{Value: "悬疑推理", Text: "悬疑推理"},
					{Value: "战争军旅", Text: "战争军旅"},
				},
				Required: true, ErrorMessage: "请选择写作背景",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "例如：增加XXX故事的深度、人物的刻画...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入内容要求",
			},
		},
	}
}

func (t *toolWeiXiaoShuo) GetMessage() (result string) {
	return t.Message
}

func (t *toolWeiXiaoShuo) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolWeiXiaoShuo) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
