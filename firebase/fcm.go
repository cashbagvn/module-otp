package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"
	"github.com/thoas/go-funk"
)

// Result ...
type Result struct {
	SuccessCount int
	FailureCount int
	ErrorTokens  []string
}

// SendToListDevices ...
func SendToListDevices(ctx context.Context, tokens []string, payload *messaging.Message) (res Result, err error) {
	var (
		errTokens []string
	)

	// remove empty fcm token before send
	tokens = removeEmptyStrings(tokens)

	for {
		// separate list tokens if exceeded limit
		send, rest := processTokens(tokens, limit)
		if len(send) <= 0 {
			break
		}
		// assign info
		multiCastMsg := &messaging.MulticastMessage{
			Tokens:       send,
			Data:         payload.Data,
			Notification: payload.Notification,
			Android:      payload.Android,
		}
		// send
		br, err := clientMessage.SendMulticast(ctx, multiCastMsg)
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return res, err
		}
		fmt.Println("Total Success : ", br.SuccessCount)
		fmt.Println("Total Failed : ", br.FailureCount)

		res.SuccessCount += br.SuccessCount
		res.FailureCount += br.FailureCount
		tokens = rest

		// append list error tokens
		errTokens = append(errTokens, getErrorTokensFromMulticastMsg(br, send)...)
	}

	return
}

func getErrorTokensFromMulticastMsg(br *messaging.BatchResponse, inputTokens []string) (tokens []string) {
	for i, r := range br.Responses {
		if !r.Success {
			tokens = append(tokens, inputTokens[i])
		}
	}
	return
}

// separate tokens for multiple time send
func processTokens(tokens []string, limit int) (send, rest []string) {
	send = tokens
	if len(tokens) > limit {
		send = tokens[:limit]
		rest = tokens[limit:]
	}
	return
}

func removeEmptyStrings(s []string) []string {
	result := funk.FilterString(s, func(item string) bool {
		return item != ""
	})
	return result
}
