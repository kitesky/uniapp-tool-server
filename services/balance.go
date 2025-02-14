package services

import (
	"app-api/dao"
	"app-api/types/schemas"

	"github.com/jinzhu/copier"
)

type balance struct{}

func NewBalance() *balance {
	return &balance{}
}

func (o *balance) GetBalanceList(req schemas.BalancePageReq) (newResult *schemas.BalanceNewPageRes, err error) {
	result, err := dao.NewBalance().GetBalanceList(req)
	newResult = &schemas.BalanceNewPageRes{}
	copier.Copy(newResult, result)
	for index, item := range newResult.Items {
		newResult.Items[index].TypeText = schemas.BalanceTypeOptions[item.Type]
		newResult.Items[index].TypeColor = schemas.BalanceTypeColorOptions[item.Type]
	}

	return
}
