package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolYaoQingHan struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolYaoQingHan() *toolYaoQingHan {
	return &toolYaoQingHan{
		Message: "写一份邀请函,邀请主题:{title},邀请对象:{to},邀请人:{from},写作风格:{style},字数限制:{size},信息补充:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "邀请主题", Placeholder: "例如：公司年会、产品发布会...",
				Required: true, ErrorMessage: "请输入邀请主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "to", Title: "邀请对象", Placeholder: "受邀人信息、名称、职位...",
				Required: true, ErrorMessage: "请输入邀请对象",
			},
			{
				FieldType: "input", ValueType: "string", Name: "from", Title: "邀请人", Placeholder: "邀请人，例如：xxx公司、xxx部门、xxx人...",
				Required: true, ErrorMessage: "请输入邀请人",
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
				Placeholder: "邀请相关信息。例如：时间、地点、联系人、会议主题内容...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入写作要求",
			},
		},
	}
}

func (t *toolYaoQingHan) GetMessage() (result string) {
	return t.Message
}

func (t *toolYaoQingHan) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolYaoQingHan) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
