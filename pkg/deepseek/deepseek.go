package deepseek

import (
	"app-api/types/consts"
	"app-api/types/schemas"
	"app-api/utils"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type deepSeek struct{}

func NewDeepSeek() *deepSeek {
	return &deepSeek{}
}

func (d *deepSeek) SendMessage(request *schemas.DeepSeekReq) (result *schemas.DeepSeekRes, err error) {
	result = &schemas.DeepSeekRes{}
	jsonStr, _ := json.Marshal(request)
	fmt.Printf("request: %+v", string(jsonStr))

	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+consts.DeepSeekApiKey).
		SetBody(string(jsonStr)).
		SetResult(result). // or SetResult(AuthSuccess{}).
		Post(consts.DeepSeekBaseUrl + "/chat/completions")

	utils.ZapLog().Info("deepseek", "SendMessage", zap.Any("request", request), zap.Any("response", resp))
	return
}
