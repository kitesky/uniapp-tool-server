package services

import (
	"app-api/dao"
	"app-api/types/schemas"

	"github.com/jinzhu/copier"
)

type reward struct{}

func NewReward() *reward {
	return &reward{}
}

func (s *reward) GetRewardList(req schemas.RewardPageReq) (newResult *schemas.RewardNewPageRes, err error) {
	result, err := dao.NewReward().GetRewardList(req)
	newResult = &schemas.RewardNewPageRes{}
	copier.Copy(newResult, result)
	for index, item := range newResult.Items {
		newResult.Items[index].TypeText = schemas.RewardTypeOptions[item.Type]
		newResult.Items[index].TypeColor = schemas.RewardTypeColorOptions[item.Type]
	}

	return
}
