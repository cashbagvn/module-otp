# module-otp

## Introduction
**[Firebase](#firebase)**<br>
**[Mailer](#mailer)**<br>
**[Zalo](#sms-zalo)**<br>
**[Vietguys](#sms-vietguys)**<br>
**[Model Sms](#support-model-sms-mongodb)**<br>
**[DAO Function](#support-function-dao)**<br>


## Firebase
### Import
`import "github.com/cashbagvn/module-otp/firebase"`
### Config
- Config information:
```
firebase.InitApp(firebase.AppConfig{
		Credentials:              <GoogleCredentialsJSONEncoded>, // base64
		ProjectID:                <GoogleFirebaseProjectID>,
		LimitFCMTokensPerMessage: 100,
})
```
### Function
- **SendToListDevices** : Send Message to devices.
```
func SendToListDevices(ctx context.Context, tokens []string, payload *messaging.Message) (res Result, err error)
```

- **VerifyIDToken** : Verify token fcm
```
func VerifyIDToken(idToken string) (*auth.Token, error)
```
Struct `*auth.Token` :
```
type Token struct {
	AuthTime int64                  json:"auth_time"
	Issuer   string                 json:"iss"
	Audience string                 json:"aud"
	Expires  int64                  json:"exp"
	IssuedAt int64                  json:"iat"
	Subject  string                 json:"sub,omitempty"
	UID      string                 json:"uid,omitempty"
	Firebase FirebaseInfo           json:"firebase"
	Claims   map[string]interface{} json:"-"
}
```

- **GetUserInfoByUID** : Verify token fcm
```
func GetUserInfoByUID(uid string) (*auth.UserRecord, error)
```
Struct `auth.UserRecord`:
```
type UserRecord struct {
	*UserInfo
	CustomClaims           map[string]interface{}
	Disabled               bool
	EmailVerified          bool
	ProviderUserInfo       []*UserInfo
	TokensValidAfterMillis int64 // milliseconds since epoch.
	UserMetadata           *UserMetadata
	TenantID               string
}
```

## Mailer
### Import
`import "github.com/cashbagvn/module-otp/mailer"`
### Config
- Struct `MailConfig` :
```
type MailConfigs struct {
	Address    string
	Port       int
	Username   string
	Password   string
	FromHeader string
}
```
- Config information:
```
var (
    configs = mailer.MailConfigs({...})
)
// Init Mailer
mailer.Init(configs)
```

### Function
- **SendMail** : Send email to mails
```
func SendMail(email Mail) (err error)
```
Struct `Mail` :
```
type Mail struct {
	Subject string
	To      []string
	CC      []string
	Body    string
}
```

## SMS Zalo
### Import
`import "github.com/cashbagvn/module-otp/zalo"`

### Config
- Struct `ConfigZalo` :
```
ConfigZalo struct {
    AccessToken string
    Host        string
    TemplateID  string
}
```
- Init:
```
zalo.Init(zalo.ConfigZalo{
    AccessToken: <AccessToken>,
    Host:        <Host>,
    TemplateID:  <TemplateID>,
})
```

### Function
- **SendOTP** : Send message code by zalo
```
func SendOTP(phone, code string) (result *Result, jsonStr string, err error)
```
Struct `Result` and `Data`:
```
//Result
Result struct {
    Message string json:"message"
    Error   int    json:"error"
    Data    *Data  json:"data"
}
// Data ...
Data struct {
    SentTime string json:"sent_time"
    MsgID    string json:"msg_id"
}
```

## SMS Vietguys
### Import
`import "github.com/cashbagvn/module-otp/vietguys"`

### Config
- Struct `VGConfig` :
```
type VGConfig struct {
	Endpoint string
	User     string
	Pwd      string
	From     string
}
```
- Init:
```
vietguys.Init(vietguys.VGConfig{
    Endpoint: <Endpoint>,
    User:     <User>,
    Pwd:      <Password>,
    From:     <From>,
})
```

### Function
- **SendOTP** : Send message sms code by VietGuys
```
unc SendOTP(phone, content string) (success bool, result VietguysResult, jsonStr string)
```
Struct `Result` and `Data`:
```
// VietguysResult ...
type VietguysResult struct {
	Carrier   string  json:"carrier" bson:"carrier"
	Error     int     json:"error" bson:"error"
	ErrorCode int     json:"error_code" bson:"errorCode"
	MsgID     string  json:"msgId" bson:"msgId"
	Message   string  json:"message" bson:"message"
	Log       string  json:"log" bson:"log"
}
```

## Support Model SMS (MongoDB)
### Import
`import "github.com/cashbagvn/module-otp/otpmodel"`
### Struct
```
// SMSLogRaw tracking request otp
type SMSLogRaw struct {
	ID        primitive.ObjectID bson:"_id"
	Service   string             bson:"service" // vietguys, ...
	Carrier   string             bson:"carrier" // mobifone, viettel, vinaphone, ...
	Type      string             bson:"type"    // otp, ...
	Recipient string             bson:"recipient"
	Content   string             bson:"content"
	CreatedAt time.Time          bson:"createdAt"
	Success   bool               bson:"success"
	Result    string             bson:"result"
}

// OtpRaw save info otp code for verify. Remove when verify success
type OtpRaw struct {
	ID        primitive.ObjectID bson:"_id"
	To        []string           bson:"to"
	ExpireAt  time.Time          bson:"expireAt"
	Code      string             bson:"code"
	Type      string             bson:"type"
	CreatedAt time.Time          bson:"createdAt"
	Tracker   string             bson:"tracker"
}

// VerifyOTPRequest
type VerifyOTPRequest struct {
	Receipt string
	Tracker string
	Code    string
}
```

## Support Function DAO 
### Import
`import "github.com/cashbagvn/module-otp/sms"`
### Function DAO
- **SaveMultiLog**
```
// SaveMultiLog
func SaveMultiLog(ctx context.Context, col *mongo.Collection, smsLogs []interface{}) // smsLogs []SMSLogRaw
```

- **SaveMultiLog**
```
// SaveMultiLog
func SaveMultiLog(ctx context.Context, col *mongo.Collection, smsLogs []interface{}) // smsLogs []SMSLogRaw
```
- **NewSMSLogRaw**
  - Params:
    - `service`: Service send sms
    - `smsType`: `sms` or `firebase`
    - `recipient`: Phone number receive code
    - `content`: Content sms
    - `result`: Result send sms
    - `success`: true or false
```
func NewSMSLogRaw(service, smsType, recipient, content, result string, success bool) *otpmodel.SMSLogRaw 
```

- **NewOtpBSON**
  - Params:
    - `to`: List phone receive code
    - `otpType`: `sms` or `firebase`
    - `tracker`: `GenerateOTPTracker(...)`
```
func NewOtpBSON(to []string, otpType, tracker string) (otp otpmodel.OtpRaw)
```

- **SaveTracking**
```
func SaveTracking(ctx context.Context, col *mongo.Collection, raw otpmodel.OtpRaw) error
```

- **VerifyOtp**
```
func VerifyOtp(ctx context.Context, col *mongo.Collection, req otpmodel.VerifyOTPRequest) (isValid bool) ...
```

- **GenerateOTPTracker**
  - Params:
    - `serviceName`: Service Call OTP
    - `collection`: Name collection
    - `targetID`: ObjectID Mongodb
```
func GenerateOTPTracker(serviceName, collection string, targetID primitive.ObjectID) string
```