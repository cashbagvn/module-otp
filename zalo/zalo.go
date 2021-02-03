package zalo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	// ConfigZalo
	ConfigZalo struct {
		AccessToken string
		Host        string
		TemplateID  string
	}

	data struct {
		OTP string `json:"OTP"`
	}

	messageReq struct {
		Phone        string `json:"phone"`
		TemplateID   string `json:"template_id"`
		TemplateData data   `json:"template_data"`
	}
	//Result
	Result struct {
		Message string `json:"message"`
		Error   int    `json:"error"`
		Data    *Data  `json:"data"`
	}
	// Data ...
	Data struct {
		SentTime string `json:"sent_time"`
		MsgID    string `json:"msg_id"`
	}
)

var (
	host        string
	templateId  string
	accessToken string
)

func Init(config ConfigZalo) {
	host = config.Host
	accessToken = config.AccessToken
	templateId = config.TemplateID
}

// SendOTP ...
func SendOTP(phone, code string) (result *Result, jsonStr string, err error) {
	payload := messageReq{
		Phone:      phone,
		TemplateID: templateId,
		TemplateData: data{
			OTP: code,
		},
	}
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, "", err
	}
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, "", err
	}
	q := req.URL.Query()
	q.Add("access_token", accessToken)

	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)

	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}
	jsonStr = string(body)
	err = json.Unmarshal(body, &result)
	return
}
