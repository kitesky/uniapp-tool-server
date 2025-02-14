package tools

import "app-api/types/schemas"

type ToolHandler interface {
	// 返回模板消息
	GetMessage() (message string)

	// 获取json schemas
	GetFormSchemas() (formItems []*schemas.ToolFormItem)

	// 请求处理
	RequestHandler(message string) (result *schemas.ToolHandlerResponse, err error)
}
