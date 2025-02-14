package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

// 测试
type toolShiTiShengCheng struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolShiTiShengCheng() *toolShiTiShengCheng {
	return &toolShiTiShengCheng{
		Message: "生成一份学生考试题,学科科目:{title},年级:{year},生成试题数量:{count},试题类型:{type},是否需要参考答案:{answer},字数限制:{size},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "title", Title: "科目", Placeholder: "输入试题的科目，例如：语文、数学、历史...",
				Required: true, ErrorMessage: "请输入科目",
			},
			{
				FieldType: "input", ValueType: "string", Name: "year", Title: "年级", Placeholder: "限定试题年级",
				Required: false, ErrorMessage: "请输入年级，例如：小学一年级",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "count", Title: "生成数量", Placeholder: "生成数量",
				Value: "1条",
				Options: []schemas.ToolFormItemSelect{
					{Value: "1条", Text: "1条"},
					{Value: "3条", Text: "3条"},
					{Value: "5条", Text: "5条"},
					{Value: "10条", Text: "10条"},
				},
				Required: true, ErrorMessage: "请选择生成条数",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "type", Title: "试题类型", Placeholder: "试题类型",
				Value: "单选题",
				Options: []schemas.ToolFormItemSelect{
					{Value: "单选题", Text: "单选题"},
					{Value: "多选题", Text: "多选题"},
					{Value: "判断题", Text: "判断题"},
					{Value: "填空题", Text: "填空题"},
					{Value: "计算题", Text: "计算题"},
					{Value: "应用题", Text: "应用题"},
					{Value: "问答题", Text: "问答题"},
				},
				Required: true, ErrorMessage: "请选择试题类型",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "answer", Title: "参考答案", Placeholder: "参考答案",
				Value: "需要",
				Options: []schemas.ToolFormItemSelect{
					{Value: "需要", Text: "需要"},
					{Value: "不需要", Text: "不需要"},
				},
				Required: true, ErrorMessage: "请选择是否需要参考答案",
			},
			{
				FieldType: "number", ValueType: "int", Name: "size", Title: "字数限制", Placeholder: "100-500字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他要求",
				Placeholder: "其他要求...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
		},
	}
}

func (t *toolShiTiShengCheng) GetMessage() (result string) {
	return t.Message
}

func (t *toolShiTiShengCheng) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolShiTiShengCheng) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
