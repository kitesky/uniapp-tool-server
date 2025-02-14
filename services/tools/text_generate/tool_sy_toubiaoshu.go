package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolTouBiaoShu struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolTouBiaoShu() *toolTouBiaoShu {
	return &toolTouBiaoShu{
		Message: "写一份投标书,投标单位:{title},全权代表:{representative},写作风格:{style},字数限制:{size},信息补充:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "投标单位", Placeholder: "投标单位、机构名称、公司全称...",
				Required: true, ErrorMessage: "请输入投标单位",
			},
			{
				FieldType: "input", ValueType: "string", Name: "representative", Title: "全权代表", Placeholder: "被授权的全权代表名称",
				Required: true, ErrorMessage: "请输入全权代表名称",
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
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "信息补充",
				Placeholder: "其他关键信息补充。例如：投标报价、设计方案、资质能力...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入其他信息",
			},
		},
	}
}

func (t *toolTouBiaoShu) GetMessage() (result string) {
	return t.Message
}

func (t *toolTouBiaoShu) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolTouBiaoShu) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
