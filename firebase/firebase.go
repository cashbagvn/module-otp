package firebase

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"log"

	firebasego "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"github.com/logrusorgru/aurora"
	"google.golang.org/api/option"
)

type AppConfig struct {
	Credentials              string
	ProjectID                string
	LimitFCMTokensPerMessage int
}

var (
	client        *auth.Client
	clientMessage *messaging.Client
	limit         int
	ctx           = context.Background()
)

func InitApp(appConfig AppConfig) {
	limit = appConfig.LimitFCMTokensPerMessage
	credentials := Base64Decode(appConfig.Credentials)
	opt := option.WithCredentialsJSON(credentials)

	cfg := &firebasego.Config{ProjectID: appConfig.ProjectID}
	app, err := firebasego.NewApp(ctx, cfg, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Init Auth
	c, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
	}

	fmt.Println(aurora.Green("*** Firebase authentication init successfully"))
	client = c
	// Init FCM
	cF, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Error getting Messaging client: %v\n", err)
	}
	fmt.Println(aurora.Green("*** FCM init successfully"))
	clientMessage = cF
}

// Base64Decode ...
func Base64Decode(text string) []byte {
	sDec, _ := b64.StdEncoding.DecodeString(text)
	return sDec
}

// Base64Encode ...
func Base64Encode(data []byte) string {
	sEnc := b64.StdEncoding.EncodeToString(data)
	return sEnc
}
