package services

import (
	"app-api/boot"
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"context"
	"net/http"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type payment struct{}

func NewPayment() *payment {
	return &payment{}
}

// 微信商家转账到用户零钱
func (s *payment) WechatTransfer(req *schemas.PaymentWechatTransferReq) (err error) {
	cfg := boot.Config.WechatPay
	client, err := wechat.NewClientV3(cfg.MchId, cfg.SerialNo, cfg.ApiV3Key, cfg.ApiclientKeyContent)
	if err != nil {
		utils.ZapLog().Error("payment", "WechatTransfer客户端错误", zap.Error(err))
		return
	}

	// 金额单位转为分
	totalAmount := decimal.NewFromFloat(req.TransferAmount).Mul(decimal.NewFromFloat(100)).RoundDown(0).IntPart()
	bm := make(gopay.BodyMap)
	bm.Set("appid", cfg.Appid).
		Set("mchid", cfg.MchId).
		Set("out_bill_no ", req.OutBillNo).
		Set("transfer_scene_id", req.TransferSceneID).
		Set("openid", req.OpenID).
		Set("transfer_amount", totalAmount).
		Set("transfer_remark", req.TransferRemark).
		// Set("notify_url", req.NotifyURL).
		Set("transfer_scene_report_infos", []schemas.PaymentWechatTransferSceneReportInfo{
			{InfoContent: "活动名称", InfoType: "新会员有礼"},
			{InfoContent: "奖励说明", InfoType: "注册会员现金奖励"},
		})
	utils.ZapLog().Info("payment", "WechatTransfer 请求数据", zap.Any("result", bm))

	resp, err := client.V3TransferBills(context.Background(), bm)
	utils.ZapLog().Info("payment", "发起V3TransactionJsapi请求响应", zap.Any("resp", resp))
	if err != nil {
		utils.ZapLog().Error("payment", "发起V3TransactionJsapi请求错误", zap.Error(err))
		return
	}

	return
}

// 微信小程序下单
func (s *payment) WechatMiniApp(req *schemas.PaymentWechatMiniAppReq) (result *schemas.PaymentWechatMiniAppRes, err error) {
	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	cfg := boot.Config.WechatPay
	client, err := wechat.NewClientV3(cfg.MchId, cfg.SerialNo, cfg.ApiV3Key, cfg.ApiclientKeyContent)
	if err != nil {
		utils.ZapLog().Error("payment", "初始化微信客户端错误", zap.Error(err))
		return
	}

	// 异步通知地址
	notifyUrl := boot.Config.App.Url + consts.WechatMiniAppPayNotifyUrl
	// 过期时间
	expireTime := time.Now().Add(consts.OrderExpireTime * time.Minute).Format(time.RFC3339)
	// 金额单位转为分
	totalAmount := decimal.NewFromFloat(req.Amount).Mul(decimal.NewFromFloat(100)).RoundDown(0).IntPart()
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", cfg.Appid).
		Set("mchid", cfg.MchId).
		Set("description", req.Description).
		Set("out_trade_no", req.OutTradeNo).
		Set("time_expire", expireTime).
		Set("notify_url", notifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", totalAmount).Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", req.OpenID)
		})

	utils.ZapLog().Info("payment", "请求参数", zap.Any("req", req), zap.Any("bm", bm))

	// 统一下单
	resp, err := client.V3TransactionJsapi(context.Background(), bm)
	utils.ZapLog().Info("payment", "发起V3TransactionJsapi请求响应", zap.Any("resp", resp))
	if err != nil {
		utils.ZapLog().Error("payment", "发起V3TransactionJsapi请求错误", zap.Error(err))
		return
	}

	// 请求小程序支付字符串
	applet, err := client.PaySignOfApplet(cfg.Appid, resp.Response.PrepayId)
	utils.ZapLog().Info("payment", "发起小程序支付请求响应", zap.Any("resp", applet))
	if err != nil {
		utils.ZapLog().Error("payment", "发起小程序支付请求响应", zap.Error(err))
		return
	}

	result = &schemas.PaymentWechatMiniAppRes{}
	copier.Copy(&result, &applet)
	return
}

// 微信支付异步通知
func (s *payment) WechatMiniAppNotify(req *http.Request) (result *schemas.PaymentWechatPayNotifyData, err error) {
	cfg := boot.Config.WechatPay
	client, err := wechat.NewClientV3(cfg.MchId, cfg.SerialNo, cfg.ApiV3Key, cfg.ApiclientKeyContent)
	if err != nil {
		utils.ZapLog().Error("payment", "异步验签初始化微信客户端错误", zap.Error(err))
		return
	}

	// 设置自动验签
	err = client.AutoVerifySign()
	if err != nil {
		utils.ZapLog().Error("payment", "启动自动验签错误", zap.Error(err))
		return
	}

	notifyReq, err := wechat.V3ParseNotify(req)
	if err != nil {
		utils.ZapLog().Error("payment", "异步验签解析回调参数错误", zap.Error(err))
		return
	}

	// 解密回调数据
	result = &schemas.PaymentWechatPayNotifyData{}
	if err = notifyReq.DecryptCipherTextToStruct(cfg.ApiV3Key, result); err != nil {
		utils.ZapLog().Error("payment", "异步验签解析回调数据错误", zap.Error(err))
		return
	}

	// 订单处理逻辑
	// 推送到订单队列中
	if _, err = NewJob().NewOrderPaidTask(&schemas.OrderPaidPayload{
		PayOrderNo: result.TransactionID,
		PayTime:    result.SuccessTime,
		OrderNo:    result.OutTradeNo,
	}); err != nil {
		utils.ZapLog().Error("payment", "订单处理逻辑错误", zap.Error(err))
		return
	}

	utils.ZapLog().Info("payment", "接收数据", zap.Any("result", result))
	return
}
