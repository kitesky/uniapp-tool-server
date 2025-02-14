package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolJiaoXueHuoDong struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolJiaoXueHuoDong() *toolJiaoXueHuoDong {
	return &toolJiaoXueHuoDong{
		Message: "写一篇教学活动设计,教学活动主题:{title},课程时长{minutes}分钟,教学风格{style},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "textarea", ValueType: "string", Name: "title", Title: "教学活动主题", Placeholder: "课堂教学活动主题，例如：复习内容、讲授新课、布置作业等",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "minutes", Title: "课程时间(分钟数)", Placeholder: "输入数字即可，例如:45",
				Required: true, ErrorMessage: "请输入课程时间",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "教学风格", Placeholder: "",
				Value:    "正式严肃",
				Options:  schemas.ToolTextGenerateStyleV3,
				Required: true, ErrorMessage: "请选择风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    1000,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他要求",
				Placeholder: "输入教学的其他要求，例如：教学的方法与方式、教学目的...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolJiaoXueHuoDong) GetMessage() (result string) {
	return t.Message
}

func (t *toolJiaoXueHuoDong) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolJiaoXueHuoDong) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
