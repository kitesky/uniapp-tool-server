package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolBanJiGuanLiZhiDu struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolBanJiGuanLiZhiDu() *toolBanJiGuanLiZhiDu {
	return &toolBanJiGuanLiZhiDu{
		Message: "写一份班级管理制度条例,制度主题:{title},学生阶段:{year},学生人数:{count},字数限制:{size}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "制度主题", Placeholder: "主题、关键词",
				Required: true, ErrorMessage: "请输入主题",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "year", Title: "学生阶段", Placeholder: "学生阶段",
				Value: "小学生",
				Options: []schemas.ToolFormItemSelect{
					{Value: "小学生", Text: "小学生"},
					{Value: "初中生", Text: "初中生"},
					{Value: "高中生", Text: "高中生"},
					{Value: "大学生", Text: "大学生"},
				},
				Required: true, ErrorMessage: "请选择学生阶段",
			},
			{
				FieldType: "number", ValueType: "int", Name: "count", Title: "学生人数", Placeholder: "学生人数",
				Value:    45,
				Required: true, ErrorMessage: "请输入学生人数",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
		},
	}
}

func (t *toolBanJiGuanLiZhiDu) GetMessage() (result string) {
	return t.Message
}

func (t *toolBanJiGuanLiZhiDu) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolBanJiGuanLiZhiDu) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
