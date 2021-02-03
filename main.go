package main

import (
	"github.com/cashbag/otp/mailer"
	"github.com/cashbag/otp/zalo"
)

func init() {
	zalo.Init(zalo.ConfigZalo{
		AccessToken: "qXluQMlCJmNK5ljhQyrC0TyT_MfgqHrMj5AjI6Y26tQn4_PQH80BRu0pnrzDvIHIfJRnEbEW4ogmCA0qJBySMV8vh6aBk559zZMSN02ITr__Fe1F4RfS9lPjY1OEwpKMyMBm91_I1Itb0xO19QD0ACq8gq5rnq1fWHsHPtgvM7IT6f55QxbuJeq1ooncj1T5dJlRLcVXCLocQUSq7linN_ikst0QXsuPfpEX6r-uGpEo6V4PKxLD1Oq1W6zArqz4jtEzPclq2WgtHjW9QiOc5fzrqm57hnySZEA4MctNJmS",
		Host:        "https://business.openapi.zalo.me/message/template",
		TemplateID:  "200393",
	})
	mailer.Init(mailer.MailConfigs{
		Address:    "smtp.gmail.com",
		Port:       587,
		Username:   "",
		Password:   "",
		FromHeader: "",
	})
}

func main() {
	zalo.SendOTP(zalo.RequestOTP{
		Phone: "84796848600",
		Code:  "123123",
	})
}
