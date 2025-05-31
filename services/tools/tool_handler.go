package tools

import (
	"app-api/dao"
	"errors"
)

type toolHandler struct{}

func NewToolHandler() *toolHandler {
	return &toolHandler{}
}

func (h *toolHandler) GetToolHandler(code string) (handler ToolHandler, err error) {
	appTool, _ := dao.NewAppTool().GetAppTool(code)
	if appTool.ID == 0 {
		err = errors.New("tool code not found")
		return nil, err
	}

	toolType := "tool:text_generate"
	switch toolType {
	case "tool:text_generate":
		handler = NewToolTextGenerateHandler(appTool)
	default:
		err = errors.New("tool type not found")
		return nil, err
	}

	return
}
