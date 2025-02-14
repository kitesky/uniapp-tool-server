package tools

import (
	"errors"
)

type toolHandler struct{}

func NewToolHandler() *toolHandler {
	return &toolHandler{}
}

func (h *toolHandler) GetToolHandler(code string) (handler ToolHandler, err error) {
	code = string(code)
	if handler, ok := ToolHandlerList[code]; ok {
		return handler, nil
	} else {
		err = errors.New("tool code not found")
	}

	return
}
