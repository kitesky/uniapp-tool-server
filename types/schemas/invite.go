package schemas

import "time"

type InviteUserList struct {
	MMID      int64     `json:"mmid"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
type InviteInfoRes struct {
	Count        int              `json:"count"`
	InviteID     int64            `json:"invite_id"`
	InviteQRCode string           `json:"invite_qrcode"`
	UserList     []InviteUserList `json:"user_list"`
	RewardList   []string         `json:"reward_list"`
}

var InviteRewardList = []string{
	"每邀请1人获得奖励5000积分",
	"受邀人在6个月内任何消费您可获得佣金奖励",
	"佣金奖励按订单实付金额的30%",
	"完成邀请任务可获得佣金加层奖励",
	"积分奖励可兑换现金奖励",
	"奖励金可随时提现",
}
