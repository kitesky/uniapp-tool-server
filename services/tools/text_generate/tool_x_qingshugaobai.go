package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolQingShuGaoBai struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolQingShuGaoBai() *toolQingShuGaoBai {
	return &toolQingShuGaoBai{
		Message: "写一篇情书告或白信，类型: {type}, 对象名称: {name}, 关系: {relation}, 内容风格: {style}, 字数限制: {size},共同经历: {future}, 未来从晶:{future}, 写作要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "radio", ValueType: "string", Name: "type", Title: "类型", Placeholder: "",
				Value: "情书",
				Options: []schemas.ToolFormItemSelect{
					{Value: "情书", Text: "情书"},
					{Value: "告白信", Text: "告白信"},
				},
				Required: true, ErrorMessage: "请选择写作类型",
			},
			{
				FieldType: "input", ValueType: "string", Name: "name", Title: "对象称呼", Placeholder: "对象的昵称或姓名，例如：亲爱的小美",
				Value:    "",
				Required: true, ErrorMessage: "请输入对象称呼",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "relation", Title: "关系", Placeholder: "",
				Value: "恋人",
				Options: []schemas.ToolFormItemSelect{
					{Value: "恋人", Text: "恋人"},
					{Value: "妻子", Text: "妻子"},
					{Value: "丈夫", Text: "丈夫"},
					{Value: "同学", Text: "同学"},
					{Value: "同事", Text: "同事"},
					{Value: "朋友", Text: "朋友"},
					{Value: "老师", Text: "老师"},
					{Value: "情人", Text: "情人"},
				},
				Required: true, ErrorMessage: "请选择关系",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
				Value: "简明",
				Options: []schemas.ToolFormItemSelect{
					{Value: "简明", Text: "简明"},
					{Value: "诗意", Text: "诗意"},
					{Value: "浪漫", Text: "浪漫"},
					{Value: "幽默", Text: "幽默"},
					{Value: "正式", Text: "正式"},
					{Value: "伤感", Text: "伤感"},
				},
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "past", Title: "共同经历",
				Placeholder: "对过往的回忆，点点滴滴",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "future", Title: "未来憧憬",
				Placeholder: "对于未来的期待",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
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

func (t *toolQingShuGaoBai) GetMessage() (result string) {
	return t.Message
}

func (t *toolQingShuGaoBai) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolQingShuGaoBai) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
