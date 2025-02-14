package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolDiaoChaWenJuan struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolDiaoChaWenJuan() *toolDiaoChaWenJuan {
	return &toolDiaoChaWenJuan{
		Message: "写一份调查问卷,问卷主题:{title},问卷类型:{category},写作风格:{style},字数限制:{size},写作要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "问卷主题", Placeholder: "问卷主题、关键词...",
				Required: true, ErrorMessage: "请输入问卷主题",
			},
			{
				FieldType: "select", ValueType: "string", Name: "category", Title: "问卷类型", Placeholder: "问卷类型",
				Value: "市场调查",
				Options: []schemas.ToolFormItemSelect{
					{Value: "市场调查", Text: "市场调查"},
					{Value: "客户满意度调查", Text: "客户满意度调查"},
					{Value: "员工满意度调查", Text: "员工满意度调查"},
					{Value: "工作满意度调查问卷", Text: "工作满意度调查问卷"},
					{Value: "服务满意度调查问卷", Text: "服务满意度调查问卷"},
					{Value: "产品调查问卷", Text: "产品调查问卷"},
					{Value: "房地产市场调查问卷", Text: "房地产市场调查问卷"},
					{Value: "物业满意度调查", Text: "物业满意度调查"},
					{Value: "客户需求调查问卷", Text: "客户需求调查问卷"},
					{Value: "品牌知名度调查问卷", Text: "品牌知名度调查问卷"},
					{Value: "市场价格调查", Text: "市场价格调查"},
					{Value: "广告效果调查问卷", Text: "广告效果调查问卷"},
					{Value: "员工职业发展调查", Text: "员工职业发展调查"},
					{Value: "员工福利调查", Text: "员工福利调查"},
					{Value: "品牌竞争力调查", Text: "品牌竞争力调查"},
					{Value: "售后调查问卷", Text: "售后调查问卷"},
					{Value: "品牌形象调查", Text: "品牌形象调查"},
				},
				Required: true, ErrorMessage: "请选择生成条数",
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
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他信息",
				Placeholder: "补充问卷关键信息...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入其他信息",
			},
		},
	}
}

func (t *toolDiaoChaWenJuan) GetMessage() (result string) {
	return t.Message
}

func (t *toolDiaoChaWenJuan) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolDiaoChaWenJuan) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
