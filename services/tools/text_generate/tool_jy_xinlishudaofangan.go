package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolXinLiShuDaoFangAn struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolXinLiShuDaoFangAn() *toolXinLiShuDaoFangAn {
	return &toolXinLiShuDaoFangAn{
		Message: "写一篇心理疏导方案,心理问题:{title},疏导对象:{content},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "textarea", ValueType: "string", Name: "title", Title: "心理问题",
				Placeholder: "描述学生的心理问题。例如：抑郁、早恋、学习焦虑症...",
				Required:    true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "input", ValueType: "string", Name: "content", Title: "疏导对象", Placeholder: "被疏导的群体对象，例如：学生、家长、老师...",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他要求",
				Placeholder: "对于心理疏导方案的其他要求。",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolXinLiShuDaoFangAn) GetMessage() (result string) {
	return t.Message
}

func (t *toolXinLiShuDaoFangAn) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolXinLiShuDaoFangAn) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
