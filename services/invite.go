package services

import (
	"app-api/dao"
	"app-api/types/schemas"
)

type invite struct{}

func NewInvite() *invite {
	return &invite{}
}

func (s *invite) GetMyInvite(userID int64) (result *schemas.InviteInfoRes, err error) {
	user, err := NewUser().GetUser(userID)
	if err != nil {
		return nil, err
	}

	inviteList, err := dao.NewInvite().GetInviteList(userID)
	if err != nil {
		return nil, err
	}

	newInviteList := []schemas.InviteUserList{}
	for _, invite := range inviteList {
		newInviteList = append(newInviteList, schemas.InviteUserList{
			MMID:      invite.User.MMID,
			Avatar:    invite.User.Avatar,
			Name:      invite.User.Name,
			CreatedAt: invite.CreatedAt,
		})
	}

	result = &schemas.InviteInfoRes{
		Count:        len(inviteList),
		InviteID:     user.Meta.InviteID,
		InviteQRCode: user.Meta.InviteQrcode,
		RewardList:   schemas.InviteRewardList,
		UserList:     newInviteList,
	}

	return
}
