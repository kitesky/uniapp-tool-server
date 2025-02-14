package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolLvXingJiHua struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolLvXingJiHua() *toolLvXingJiHua {
	return &toolLvXingJiHua{
		Message: "制作一份旅行计划,目的地:{to},出发地:{from},旅行天数:{days},出行交通:{type},兴趣爱好:{interest},其他要求:{requirement}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "to", Title: "目的地", Placeholder: "出行活动的目的地",
				Value:    "",
				Required: true, ErrorMessage: "请输入目的地",
			},
			{
				FieldType: "input", ValueType: "string", Name: "from", Title: "出发地", Placeholder: "出发地",
				Value:    "",
				Required: true, ErrorMessage: "请输入出发地",
			},
			{
				FieldType: "number", ValueType: "int", Name: "days", Title: "旅行天数", Placeholder: "旅行天数",
				Value:    1,
				Required: true, ErrorMessage: "请输入旅行天数",
			},
			{
				FieldType: "radio", ValueType: "string", Name: "type", Title: "出行方式", Placeholder: "",
				Value: "自驾",
				Options: []schemas.ToolFormItemSelect{
					{Value: "自驾", Text: "自驾"},
					{Value: "飞机", Text: "飞机"},
					{Value: "火车", Text: "火车"},
					{Value: "客车", Text: "客车"},
					{Value: "骑行", Text: "骑行"},
					{Value: "徒步", Text: "徒步"},
					{Value: "顺风车", Text: "顺风车"},
					{Value: "其他", Text: "其他"},
				},
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "interest", Title: "兴趣爱好",
				Placeholder: "例如：历史建筑、自然风景、人文地理、美食...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入要求",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "partner", Title: "同行人员",
				Placeholder: "一起出行的人员、人数描述、兴趣爱好描述...",
				Value:       "",
				Required:    false, ErrorMessage: "请输入同行人员",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "requirement", Title: "其他要求",
				Placeholder: "其他需求信息描述",
				Value:       "",
				Required:    false, ErrorMessage: "请输入内容要求",
			},
		},
	}
}

func (t *toolLvXingJiHua) GetMessage() (result string) {
	return t.Message
}

func (t *toolLvXingJiHua) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolLvXingJiHua) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
