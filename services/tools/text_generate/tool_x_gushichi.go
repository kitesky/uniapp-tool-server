package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolGuShiCi struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolGuShiCi() *toolGuShiCi {
	return &toolGuShiCi{
		Message: "写一篇古诗词，题目: {title}, 类型: {type}, 辩论风格: {style}, 字数限制: {size}, 其他要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "题目", Placeholder: "题目",
				Value:    "",
				Required: true, ErrorMessage: "请输入题目",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "type", Title: "诗词类型", Placeholder: "",
				Value: "诗",
				Options: []schemas.ToolFormItemSelect{
					{Value: "诗", Text: "诗"},
					{Value: "词", Text: "词"},
				},
				Required: true, ErrorMessage: "请选择诗词类型",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "诗词风格", Placeholder: "",
				Value: "辞藻华丽",
				Options: []schemas.ToolFormItemSelect{
					{Value: "辞藻华丽", Text: "辞藻华丽"},
					{Value: "雅致清新", Text: "雅致清新"},
					{Value: "明艳欢快", Text: "明艳欢快"},
					{Value: "质朴平实", Text: "质朴平实"},
					{Value: "温柔浪漫", Text: "温柔浪漫"},
					{Value: "伤感忧虑", Text: "伤感忧虑"},
					{Value: "悲凉凄清", Text: "悲凉凄清"},
					{Value: "干净治愈", Text: "干净治愈"},
					{Value: "激情昂扬", Text: "激情昂扬"},
					{Value: "雄浑开阔", Text: "雄浑开阔"},
					{Value: "感恩怀念", Text: "感恩怀念"},
				},
				Required: true, ErrorMessage: "请选择诗词风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "topic", Title: "诗词形式", Placeholder: "五言绝句、七言律诗...",
				Value:    "",
				Required: true, ErrorMessage: "请输入诗词形式",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "创作要求",
				Placeholder: "例如：借景抒情、托物言志、忧国忧民...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolGuShiCi) GetMessage() (result string) {
	return t.Message
}

func (t *toolGuShiCi) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolGuShiCi) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
