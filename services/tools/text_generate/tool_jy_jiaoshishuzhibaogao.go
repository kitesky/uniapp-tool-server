package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolJiaoShiShuZhiBaoGao struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolJiaoShiShuZhiBaoGao() *toolJiaoShiShuZhiBaoGao {
	return &toolJiaoShiShuZhiBaoGao{
		Message: "写一篇教师述职报告,学科职务:{title},语言风格:{style},字数限制:{size},其他信息:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "textarea", ValueType: "string", Name: "title", Title: "学科职务", Placeholder: "学科职务",
				Required: true, ErrorMessage: "请输入学科职务",
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
				Placeholder: "补充其他工作相关信息，例如：工作内容、工作职责、取得成果、未来计划...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入其他信息",
			},
		},
	}
}

func (t *toolJiaoShiShuZhiBaoGao) GetMessage() (result string) {
	return t.Message
}

func (t *toolJiaoShiShuZhiBaoGao) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolJiaoShiShuZhiBaoGao) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
