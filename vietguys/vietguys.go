package vietguys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type VGConfig struct {
	Endpoint string
	User     string
	Pwd      string
	From     string
}

type Request struct {
	Receipts []string
	Content  string
	Tracker  string
}

// VietguysResult ...
type VietguysResult struct {
	Carrier   string `json:"carrier" bson:"carrier"`
	Error     int    `json:"error" bson:"error"`
	ErrorCode int    `json:"error_code" bson:"errorCode"`
	MsgID     string `json:"msgId" bson:"msgId"`
	Message   string `json:"message" bson:"message"`
	Log       string `json:"log" bson:"log"`
}

var (
	config VGConfig
)

// Init
func Init(cfg VGConfig) {
	config = cfg
}

// SendOTP ...
func SendOTP(phone, content string) (success bool, result VietguysResult, jsonStr string) {
	// Create payload
	params := url.Values{}
	params.Add("u", config.User)
	params.Add("pwd", config.Pwd)
	params.Add("from", config.From)
	params.Add("json", "1")
	params.Add("phone", phone)
	params.Add("sms", content)
	payload := strings.NewReader(params.Encode())

	// Create request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, config.Endpoint, payload)
	if err != nil {
		return
	}

	// Add necessary headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	if result.Error == 0 {
		success = true
	}

	return
}
