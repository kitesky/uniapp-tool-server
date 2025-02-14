package services

import (
	"app-api/boot"
	"app-api/dao"
	"app-api/models"
	"app-api/services/tools"
	"encoding/json"
	"fmt"
	"strings"

	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type tool struct{}

func NewTool() *tool {
	return &tool{}
}

// 工具使用
func (s *tool) HandleToolUsedTask(req *schemas.TaskToolUsedPayload) (err error) {
	// 完成工具体验任务
	NewJob().NewTaskCompletedTask(&schemas.TaskCompletedPayload{
		TaskCode: consts.TaskToolUsedTextOnce,
		UserID:   req.UserID,
	})

	return
}

// 构造消息
func (s *tool) GetToolData(handler tools.ToolHandler, jsonString string) (message string, schemas string, err error) {
	// 验证是否json
	if !gjson.Valid(jsonString) {
		err = errors.New("无效的JSON数据")
		return
	}

	// 替换模板中的变量值
	message = handler.GetMessage()
	// 获取表单字段
	formSchemas := handler.GetFormSchemas()

	// 遍历赋值
	for index, schema := range formSchemas {
		val := gjson.Get(jsonString, schema.Name)
		switch val.Type {
		case gjson.String:
			formSchemas[index].Value = val.String()
			message = strings.ReplaceAll(message, fmt.Sprintf("{%s}", schema.Name), val.String())
		case gjson.Number:
			formSchemas[index].Value = val.Int()
			message = strings.ReplaceAll(message, fmt.Sprintf("{%s}", schema.Name), val.String())
		}
	}

	// 格式化消息
	schemasObj, err := json.Marshal(formSchemas)
	if err != nil {
		err = errors.New("json解析失败")
		return
	}
	schemas = string(schemasObj)
	return
}

// 工具请求处理
func (s *tool) ToolHandler(userID int64, code string, jsonParams string) (result *schemas.ToolHandlerResponse, err error) {
	utils.ZapLog().Info("tool", "收到请求", zap.Any("userID", userID), zap.Any("code", code), zap.Any("jsonParams", jsonParams))

	// tool 信息
	tool, err := dao.NewTool().GetToolWithCode(code)
	if err != nil {
		utils.ZapLog().Info("tool", "获取tool信息错误", zap.Error(err))
		err = errors.New("工具应用已下架或信息错误")
		return
	}

	// 获取操作句柄
	handler, err := tools.NewToolHandler().GetToolHandler(code)
	if err != nil {
		utils.ZapLog().Info("tool", "获取handler句柄错误[GetToolHandler]", zap.Error(err))
		err = errors.New("获取工具错误")
		return
	}

	// 获取数据
	message, formSchemas, err := s.GetToolData(handler, jsonParams)
	if err != nil {
		utils.ZapLog().Info("tool", "获取数据失败[GetToolData]", zap.Error(err))
		err = errors.New("获取数据失败")
		return
	}

	// 提交日志
	tx := boot.DB.Begin()
	logRes, err := dao.NewActivity().ActivityLog(tx, &models.UserActivityLog{
		UserID:       userID,
		UUID:         utils.GenerateStringUniqueID(),
		Code:         code,
		Title:        tool.Title,
		Description:  tool.Description,
		Amount:       tool.Price,
		RequestBody:  jsonParams,
		FormSchemas:  formSchemas,
		ContentType:  "",
		Content:      "",
		ResponseBody: "",
	})

	if err != nil {
		utils.ZapLog().Info("tool", "ActivityLog记录日志错误", zap.Error(err))
		tx.Rollback()
		return
	}

	// 处理请求
	result, err = handler.RequestHandler(message)
	if err != nil {
		utils.ZapLog().Info("tool", "处理请求失败", zap.Error(err))
		tx.Commit()
		return
	}

	// 变更结果
	result.UUID = logRes.UUID
	if result.Status == "success" {
		logRes.Status = "success"
		logRes.ContentType = result.ContentType
		logRes.Content = result.Content
		logRes.ResponseBody = result.ResponseBody
		err = tx.Save(logRes).Error
		tx.Commit()

		// 完成体验任务
		NewJob().NewTaskCompletedTask(&schemas.TaskCompletedPayload{
			TaskCode: consts.TaskToolUsedTextOnce,
			UserID:   userID,
		})
		return
	}

	tx.Rollback()
	return nil, errors.New("处理请求失败")
}

func (s *tool) GetToolList(req schemas.ToolPageReq) (newResult *schemas.ToolNewPageRes, err error) {
	result, err := dao.NewTool().GetToolList(req)
	newResult = &schemas.ToolNewPageRes{}
	copier.Copy(newResult, result)

	return
}

func (s *tool) GetTool(id int64) (newResult *schemas.ToolRes, err error) {
	result, err := dao.NewTool().GetTool(id)
	if err != nil {
		utils.ZapLog().Info("tool", "获取tool信息错误", zap.Error(err))
		err = errors.New("工具应用已下架或信息错误")
		return
	}

	newResult = &schemas.ToolRes{}
	copier.Copy(newResult, result)

	// 获取操作句柄
	handler, err := tools.NewToolHandler().GetToolHandler(result.Code)
	if err != nil {
		utils.ZapLog().Info("tool", "获取handler句柄错误", zap.Error(err))
		err = errors.New("获取工具错误")
		return
	}
	newResult.FormSchemas = handler.GetFormSchemas()
	return
}

func (s *tool) GetRecommentTool() (list []*models.AppTool, err error) {
	err = boot.DB.Where("recommend", "Y").Find(&list).Error
	return list, err
}
