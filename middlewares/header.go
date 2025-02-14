package middlewares

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"time"

	"app-api/boot"
	"app-api/pkg/response"
	"app-api/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	AppClientID       string        = "9f823d76-5a8f-4ca3-94bc-c80febdd6501"
	AppClientSecret   string        = "96a9c3b06d93402a743605eedfbe72106fe1e967f1ec5f5bed7e8da90a033f1a"
	SignatureRedisKey string        = "signature:%s"
	SignatureExpire   time.Duration = 60 * time.Second
	TimestampMaxDiff  int64         = 15
)

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform tasks before the route handling function
		utils.ZapLog().Info("header", "接收请求:", zap.Any("requestHeader", c.Request.Header))

		clientID := c.Request.Header.Get("Client-ID")
		nonce := c.Request.Header.Get("Nonce")
		timestamp := c.Request.Header.Get("Timestamp")
		version := c.Request.Header.Get("Version")
		signature := c.Request.Header.Get("Signature")
		method := c.Request.Method

		if clientID != AppClientID {
			response.New(c).SetMessage("ClientID 错误").Error()
			c.Abort()
			return
		}

		if nonce == "" || timestamp == "" || signature == "" {
			response.New(c).SetMessage("缺少参数").Error()
			c.Abort()
			return
		}

		// decTimestamp, err := decimal.NewFromString(timestamp)
		// if err != nil {
		// 	response.New(c).SetMessage("时间戳格式错误").Error()
		// 	c.Abort()
		// 	return
		// }

		// var newTimestamp int64
		// if len(timestamp) == 13 {
		// 	newTimestamp = decTimestamp.Div(decimal.NewFromInt(1000)).IntPart()
		// } else {
		// 	newTimestamp = decTimestamp.IntPart()
		// }

		// // 时间戳误差允许在15秒内
		// if newTimestamp+TimestampMaxDiff < time.Now().Unix() {
		// 	utils.ZapLog().Info("header", "时间戳过期: ", zap.Any("timestamp", newTimestamp))
		// 	response.New(c).SetMessage("时间戳过期").Error()
		// 	c.Abort()
		// 	return
		// }

		// 签名过期或已使用过
		signatureUsed := boot.Redis.Get(context.Background(), fmt.Sprintf(SignatureRedisKey, signature)).Val()
		if signatureUsed != "" {
			utils.ZapLog().Info("header", "签名过期或已使用过: ", zap.Any("signature", signature))
			response.New(c).SetMessage("签名已过期").Error()
			c.Abort()
			return
		}

		boot.Redis.SetEx(context.Background(), fmt.Sprintf(SignatureRedisKey, signature), signature, SignatureExpire)

		var body string
		switch method {
		case "GET":
			body = ""
		case "POST":
			bodyRaw, _ := c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyRaw))
			contentType := c.Request.Header.Get("Content-Type")
			isJson := strings.Contains(contentType, "application/json")
			if len(bodyRaw) > 0 && isJson {
				body = base64.StdEncoding.EncodeToString(bodyRaw)
			}
		default:
			body = ""
		}

		// 拼接参数字符串
		plainText := strings.Join([]string{AppClientID, nonce, timestamp, version, body}, "")
		// 记录body日志
		utils.ZapLog().Info("header", "签名拼接字符串", zap.Any("plainText", plainText))
		// 验证签名
		makeSignature := utils.HmacSha256([]byte(plainText), []byte(AppClientSecret))
		if makeSignature != signature {
			utils.ZapLog().Info("header", "签名错误", zap.Any("makeSignature", makeSignature), zap.Any("reqSignature", signature))
			response.New(c).SetMessage("签名错误").Error()
			c.Abort()
			return
		}

		c.Next()
	}
}
