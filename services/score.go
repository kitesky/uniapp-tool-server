package services

import (
	"app-api/dao"
	"app-api/types/schemas"

	"github.com/jinzhu/copier"
)

type score struct{}

func NewScore() *score {
	return &score{}
}

func (s *score) GetScoreList(req schemas.ScorePageReq) (newResult *schemas.ScoreNewPageRes, err error) {
	result, err := dao.NewScore().GetScoreList(req)
	newResult = &schemas.ScoreNewPageRes{}
	copier.Copy(newResult, result)
	for index, item := range newResult.Items {
		newResult.Items[index].TypeText = schemas.ScoreTypeOptions[item.Type]
		newResult.Items[index].TypeColor = schemas.ScoreTypeColorOptions[item.Type]
	}

	return
}
