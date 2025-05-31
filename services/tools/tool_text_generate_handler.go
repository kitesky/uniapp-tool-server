package tools

import (
	"app-api/models"
	"app-api/pkg/deepseek"
	"app-api/types/schemas"
	"encoding/json"
)

type toolTextGenerateHandler struct {
	Message     string
	FormSchemas []*schemas.ToolFormItem
}

func NewToolTextGenerateHandler(appTool *models.AppTool) *toolTextGenerateHandler {
	var formSchemas []*schemas.ToolFormItem
	if err := json.Unmarshal([]byte(appTool.FormSchemas), &formSchemas); err != nil {
		formSchemas = []*schemas.ToolFormItem{}
	}

	return &toolTextGenerateHandler{
		Message:     appTool.Message,
		FormSchemas: formSchemas,
	}
}

func (t *toolTextGenerateHandler) GetMessage() (result string) {
	return t.Message
}

func (t *toolTextGenerateHandler) GetFormSchemas() (result []*schemas.ToolFormItem) {
	return t.FormSchemas
}

func (t *toolTextGenerateHandler) RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error) {
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
