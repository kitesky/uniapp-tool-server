package services

import (
	"app-api/dao"
	"app-api/types/schemas"

	"github.com/jinzhu/copier"
)

type activity struct{}

func NewActivity() *activity {
	return &activity{}
}

func (s *activity) GetActivityList(req schemas.ActivityPageReq) (newResult *schemas.ActivityNewPageRes, err error) {
	result, err := dao.NewActivity().GetActivityList(req)
	newResult = &schemas.ActivityNewPageRes{}
	copier.Copy(newResult, result)
	for index, item := range newResult.Items {
		newResult.Items[index].ContentTypeText = schemas.ActivityContentTypeOptions[item.ContentType].Title
		newResult.Items[index].StatusText = schemas.ActivityStatusOptions[item.Status].Text
		newResult.Items[index].StatusColor = schemas.ActivityStatusOptions[item.Status].Color
	}

	return
}
