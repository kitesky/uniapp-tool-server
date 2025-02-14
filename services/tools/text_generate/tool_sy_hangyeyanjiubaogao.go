package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolHangYeYanJiuBaoGao struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolHangYeYanJiuBaoGao() *toolHangYeYanJiuBaoGao {
	return &toolHangYeYanJiuBaoGao{
		Message: "写一份行业研究报告,行业信息:{title},全权代表:{representative},写作风格:{style},字数限制:{size},信息补充:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "行业信息", Placeholder: "行业、工作领域、角色...",
				Required: true, ErrorMessage: "请输入行业信息",
			},
			{
				FieldType: "select", ValueType: "string", Name: "style", Title: "写作风格", Placeholder: "",
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
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "写作要求",
				Placeholder: "其他要求。例如：分析市场前景、市场趋势及风险、竞品信息...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入写作要求",
			},
		},
	}
}

func (t *toolHangYeYanJiuBaoGao) GetMessage() (result string) {
	return t.Message
}

func (t *toolHangYeYanJiuBaoGao) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolHangYeYanJiuBaoGao) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
