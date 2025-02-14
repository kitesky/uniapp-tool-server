package services

import (
	"app-api/boot"
	"app-api/utils"
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type wechatApp struct{}

const (
	// 测试账号
	// WechatAppAppid  = "wxad50d3b12eb5f40c"
	// WechatAppSecret = "59c51e321dfba607d1ae9af2c2f9298f"

	// 微信API
	WechatAppTokenAPI          string = "https://api.weixin.qq.com/cgi-bin/token"
	WechatAppJscode2sessionAPI string = "https://api.weixin.qq.com/sns/jscode2session"
	WechatGetwxacodeunlimit    string = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

type WechatAppAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type WechatAppJscode2session struct {
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errmsg     string `json:"errmsg"`
	Openid     string `json:"openid"`
	Errcode    int    `json:"errcode"`
}

func NewWechatApp() *wechatApp {
	return &wechatApp{}
}

// 获取小程序码
func (s *wechatApp) GetMiniAppQRCode(scene string) (body []byte, err error) {
	token, _ := s.GetAccessToken()

	client := resty.New()
	params := map[string]string{
		"scene": scene,
		"page":  "pages/index/index",
		"width": "600",
	}
	jsonStr, _ := json.Marshal(params)
	url := WechatGetwxacodeunlimit + "?access_token=" + token.AccessToken
	resp, err := client.R().SetBody(string(jsonStr)).Post(url)
	body = resp.Body()
	return
}

func (s *wechatApp) GetAccessToken() (result *WechatAppAccessToken, err error) {
	client := resty.New()
	params := map[string]string{
		"grant_type": "client_credential",
		"appid":      boot.Config.WechatPay.Appid,
		"secret":     boot.Config.WechatPay.AppSecret,
	}

	body := &WechatAppAccessToken{}
	client.R().SetQueryParams(params).SetResult(body).Get(WechatAppTokenAPI)
	if body.AccessToken == "" {
		err = errors.New("获取AccessToken失败")
	}

	return body, err
}

func (s *wechatApp) Code2Session(code string) (result *WechatAppJscode2session, err error) {
	client := resty.New()
	params := map[string]string{
		"appid":      boot.Config.WechatPay.Appid,
		"secret":     boot.Config.WechatPay.AppSecret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	body := &WechatAppJscode2session{}

	// 微信响应类型 Content-Type:text/plain SetResult不生效
	resp, _ := client.R().SetQueryParams(params).Get(WechatAppJscode2sessionAPI)
	if err = json.Unmarshal(resp.Body(), body); err != nil {
		err = errors.New("解析数据格式错误")
	}

	// 记录日志
	utils.ZapLog().Info("wechatApp", "Code2Session", zap.Any("params", params), zap.Any("resp", body))
	if body.Openid == "" {
		err = errors.New(body.Errmsg)
	}

	return body, err
}
