package services

import (
	"app-api/dao"
	"app-api/types/schemas"

	"github.com/jinzhu/copier"
)

type point struct{}

func NewPoint() *point {
	return &point{}
}

func (s *point) GetPointList(req schemas.PointPageReq) (newResult *schemas.PointNewPageRes, err error) {
	result, err := dao.NewPoint().GetPointList(req)
	newResult = &schemas.PointNewPageRes{}
	copier.Copy(newResult, result)
	for index, item := range newResult.Items {
		newResult.Items[index].TypeText = schemas.PointTypeOptions[item.Type]
		newResult.Items[index].TypeColor = schemas.PointTypeColorOptions[item.Type]
	}

	return
}
