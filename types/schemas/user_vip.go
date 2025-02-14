package schemas

type UserInfoVIPRes struct {
	Active     string                 `json:"active"`
	ExpireTime Datetime               `json:"expire_time"`
	Benefits   []*UserInfoVIPBenefits `json:"benefits"`
}

type UserInfoVIPBenefits struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

var UserVIPBenefits = []*UserInfoVIPBenefits{
	{Title: "无限使用", Description: "全场AI免费用", Icon: "/assets/icons/vip-qy-1.svg"},
	{Title: "每月礼包", Description: "每月礼品赠送", Icon: "/assets/icons/vip-qy-2.svg"},
	{Title: "体验官", Description: "新品优先体验", Icon: "/assets/icons/vip-qy-3.svg"},
	{Title: "好友福袋", Description: "好友领取福袋现金", Icon: "/assets/icons/vip-qy-4.svg"},
}
