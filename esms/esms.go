package esms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ESmsConfig struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
	BrandName string `json:"brandName"`
}

type EsmsResult struct {
	CodeResult      string `json:"CodeResult"`
	CountRegenerate int    `json:"CountRegenerate"`
	ErrorMessage    string `json:"ErrorMessage"`
	SMSID           string `json:"SMSID"`
}

type smsReq struct {
	ApiKey    string `json:"ApiKey"`
	Content   string `json:"Content"`
	Phone     string `json:"Phone"`
	SecretKey string `json:"SecretKey"`
	IsUnicode string `json:"IsUnicode"`
	Brandname string `json:"Brandname"`
	SmsType   string `json:"SmsType"`
}

const EndPoint = "http://rest.esms.vn/MainService.svc/json/SendMultipleMessage_V4_post_json/"

var (
	config ESmsConfig
)

// Init
func Init(cfg ESmsConfig) {
	config = cfg
}

// SendOTP
func SendOTP(phone, content string) (success bool, result *EsmsResult, jsonStr string) {
	payload := smsReq{
		ApiKey:    config.ApiKey,
		Content:   content,
		Phone:     phone,
		SecretKey: config.SecretKey,
		IsUnicode: "0",
		Brandname: config.BrandName,
		SmsType:   "2",
	}
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return false, nil, ""
	}

	// Create request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, EndPoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}

	// Add necessary headers
	req.Header.Add("Content-Type", "application/json")

	// Call
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Make sure close body
	defer res.Body.Close()

	// Ready body
	body, err := ioutil.ReadAll(res.Body)
	jsonStr = string(body)
	fmt.Println(jsonStr)

	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}

	if result.CodeResult == "100" {
		success = true
	}

	return
}
