package text_generate

import (
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolNianZhongZongJie struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolNianZhongZongJie() *toolNianZhongZongJie {
	return &toolNianZhongZongJie{
		Message: "写一篇年终总结，行业部门: {hy}, 职业: {zw}, 写作风格: {fg}, 字数限制: {zs}, 工作描述: {ms}",
		FormSchemas: []*schemas.ToolFormItem{
			{
				FieldType: "input", ValueType: "string", Name: "hy", Title: "行业部门", Placeholder: "你所在的行业/单位/部门",
				Value:    "",
				Required: true, ErrorMessage: "请输入行业部门",
			},
			{
				FieldType: "input", ValueType: "string", Name: "zw", Title: "职位", Placeholder: "例如: 前端工程师/产品经理/销售员",
				Value:    "",
				Required: true, ErrorMessage: "请输入职位",
			},
			{
				FieldType: "select", ValueType: "string", Name: "fg", Title: "写作风格", Placeholder: "",
				Value:    "正式严谨",
				Options:  schemas.ToolTextGenerateStyleV1,
				Required: true, ErrorMessage: "请选择写作风格",
			},
			{
				FieldType: "number", ValueType: "int", Name: "zs", Title: "字数限制", Placeholder: "100-5000字",
				Value:    500,
				Required: true, ErrorMessage: "请输入字数限制",
			},
			{
				FieldType: "textarea", ValueType: "string", Name: "ms", Title: "工作描述", Placeholder: "工作内容/主要任务/业绩情况等，越详细越好。",
				Value:    "",
				Required: true, ErrorMessage: "请输入工作描述",
			},
		},
	}
}

func (t *toolNianZhongZongJie) GetMessage() (result string) {
	return t.Message
}

func (t *toolNianZhongZongJie) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolNianZhongZongJie) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
