package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolJiaoYanGongZuoZongJie struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolJiaoYanGongZuoZongJie() *toolJiaoYanGongZuoZongJie {
	return &toolJiaoYanGongZuoZongJie{
		Message: "写一篇教研工作总结,教研主题:{title},语言风格:{style},字数限制:{size},其他信息:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "textarea", ValueType: "string", Name: "title", Title: "教研主题", Placeholder: "教研主题",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "语言风格", Placeholder: "",
				Value:    "正式",
				Options:  schemas.ToolTextGenerateStyleV2,
				Required: true, ErrorMessage: "请选择风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他信息",
				Placeholder: "补充培训活动的细节信息，例如：教学方法、教学目的、本次活动的感受...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入其他信息",
			},
		},
	}
}

func (t *toolJiaoYanGongZuoZongJie) GetMessage() (result string) {
	return t.Message
}

func (t *toolJiaoYanGongZuoZongJie) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolJiaoYanGongZuoZongJie) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
