package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolXianDaiShi struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolXianDaiShi() *toolXianDaiShi {
	return &toolXianDaiShi{
		Message: "写一篇现代诗，写作主题: {title}, 写作风格: {style}, 诗歌类型: {type}, 字数限制: {size}, 写作要求: {requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "写作主题", Placeholder: "诗歌主题",
				Value:    "",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
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
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "input", ValueType: "string", Name: "type", Title: "诗歌类型",
				Placeholder: "例如：抒情诗、爱情诗、古风诗...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "其他细节要求，越详细诗歌质量越高",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolXianDaiShi) GetMessage() (result string) {
	return t.Message
}

func (t *toolXianDaiShi) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolXianDaiShi) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
