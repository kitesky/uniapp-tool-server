package schemas

import "time"

type PaymentWechatTransferSceneReportInfo struct {
	InfoType    string `json:"info_type" binding:"required"`    //信息类型 如：转账场景为1000-现金营销，需填入活动名称、奖励说明
	InfoContent string `json:"info_content" binding:"required"` //信息内容 如新会员有礼
}

type PaymentWechatTransferReq struct {
	TransferAmount  float64 `json:"transfer_amount" binding:"required"`   // 订单金额
	OutBillNo       string  `json:"out_bill_no" binding:"required"`       // 订单号
	TransferSceneID string  `json:"transfer_scene_id" binding:"required"` // 转账场景
	OpenID          string  `json:"openid" binding:"required"`            // 支付者openid
	TransferRemark  string  `json:"transfer_remark" binding:"required"`   // 转账备注
	NotifyURL       string  `json:"notify_url"`                           // 转账备注
	// UserRecvPerception       string                                 `json:"user_recv_perception"`                           // 用户收款感知
	// TransferSceneReportInfos []PaymentWechatTransferSceneReportInfo `json:"transfer_scene_report_infos" binding:"required"` // 转账备注
}

type PaymentWechatMiniAppReq struct {
	Amount      float64 `json:"amount"`       // 订单金额
	OutTradeNo  string  `json:"out_trade_no"` // 订单号
	Description string  `json:"description"`  // 订单描述
	OpenID      string  `json:"openid"`       // 支付者openid
}

type PaymentWechatMiniAppRes struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type PaymentWechatPayCallbackRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaymentWechatPayNotifyData struct {
	TransactionID string    `json:"transaction_id"`
	TradeState    string    `json:"trade_state"`
	SuccessTime   time.Time `json:"success_time"`
	OutTradeNo    string    `json:"out_trade_no"`
}
