package services

import (
	"app-api/utils"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	api             string = "https://www.idcd.com/api"
	clientID        string = "9256f951-9559-4f20-8346-ccc1fcf6a3bb"
	clientSecret    string = "721b32ecbdeb0c0656050f93ec1746b415fffdffc52ffe11597adb37d009e09c"
	signatureMethod string = "HmacSHA256"
)

type headers struct {
	ClientID        string `json:"ClientID"`
	SignatureMethod string `json:"SignatureMethod"`
	Nonce           string `json:"Nonce"`
	Timestamp       string `json:"Timestamp"`
	Signature       string `json:"Signature"`
}

type openAPI struct{}

func NewOpenAPI() *openAPI {
	return &openAPI{}
}

func (s *openAPI) Whois(domain string) (any, error) {
	url := api + "/domain/whois"

	return url, nil
}

func (s *openAPI) Get(apiURL string, query map[string]interface{}) (interface{}, error) {
	params := url.Values{}
	rawUrl, _ := url.Parse(apiURL)
	for k, v := range query {
		params.Set(k, v.(string))
	}
	rawUrl.RawQuery = params.Encode()
	newURL := rawUrl.String()

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", newURL, nil)
	if err != nil {
		return nil, err
	}

	headers := s.MakeHeader()
	req.Header.Set("ClientID", headers.ClientID)
	req.Header.Set("SignatureMethod", headers.SignatureMethod)
	req.Header.Set("Nonce", headers.Nonce)
	req.Header.Set("Timestamp", headers.Timestamp)
	req.Header.Set("Signature", headers.Signature)

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // 确保关闭响应体

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonData := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	if jsonData["code"].(float64) != 200 {
		return nil, errors.New(jsonData["message"].(string))
	}

	return jsonData["data"], nil
}

func (s *openAPI) Post(apiURL string, bodyMap map[string]interface{}) (interface{}, error) {
	// 创建HTTP客户端
	client := &http.Client{}

	jsonStr, _ := json.Marshal(bodyMap)

	// 创建HTTP请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	headers := s.MakeHeader()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ClientID", headers.ClientID)
	req.Header.Set("SignatureMethod", headers.SignatureMethod)
	req.Header.Set("Nonce", headers.Nonce)
	req.Header.Set("Timestamp", headers.Timestamp)
	req.Header.Set("Signature", headers.Signature)

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // 确保关闭响应体

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonData := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	if jsonData["code"].(float64) != 200 {
		return nil, errors.New(jsonData["message"].(string))
	}

	return jsonData["data"], nil
}

func (s *openAPI) MakeHeader() headers {
	nonce := utils.GenerateRandomString(16)
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)
	plainText := clientID + nonce + timestampStr + signatureMethod
	signature := utils.HmacSha256([]byte(plainText), []byte(clientSecret))

	return headers{
		ClientID:        clientID,
		SignatureMethod: signatureMethod,
		Nonce:           nonce,
		Timestamp:       timestampStr,
		Signature:       signature,
	}
}
