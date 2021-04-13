package firebase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"firebase.google.com/go/messaging"
	"github.com/thoas/go-funk"
)

const (
	errTopicNotAllowed = "topic not allowed"
	errEmptyTokens = "Empty tokens"
)

var (
	topics = []string{
		"all",
		"android",
		"ios",
	}
	LimitTokensForSubscribe = 1000
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

// SendWithCombineTopics ...
func SendWithCombineTopics(topics []string, msg *messaging.Message) {
	if len(topics) <= 0 {
		return
	}
	var topicCond []string

	for _, topic := range topics {
		if !IsAllowedTopic(topic) {
			continue
		}
		cond := fmt.Sprintf("'%s' in topics", topic)
		topicCond = append(topicCond, cond)
	}

	msg.Condition = strings.Join(topicCond, " && ")

	_, err := clientMessage.Send(context.Background(), msg)
	if err != nil {
		fmt.Printf("Error send message with condition %s: %v", msg.Condition, err)
	}
}


// SubscribeTokensToTopic ...
func SubscribeTokensToTopic(topic string, tokens []string) (r Result, err error) {
	if len(tokens) <= 0 {
		return
	}
	ctx := context.Background()

	tokens = removeEmptyStrings(tokens)
	if !IsAllowedTopic(topic) {
		err = errors.New(errTopicNotAllowed)
		return
	}

	for {
		// separate list tokens if exceeded limit
		send, rest := processTokens(tokens, LimitTokensForSubscribe)
		if len(send) <= 0 {
			break
		}

		res, err := clientMessage.SubscribeToTopic(ctx, tokens, topic)
		if err != nil {
			fmt.Printf("Subscribe to topic %s error: %v", topic, err)
			return r, err
		}
		r.SuccessCount += res.SuccessCount
		r.FailureCount += res.FailureCount
		fmt.Printf("Subscribe tokens to topic %s: success %d, failed %d \n", topic, res.SuccessCount, res.FailureCount)

		// get list error tokens
		if len(res.Errors) > 0 {
			r.ErrorTokens = append(r.ErrorTokens, getErrTokensFromSubscribe(res, send)...)
		}

		tokens = rest
	}

	return r, nil
}

func getErrTokensFromSubscribe(r *messaging.TopicManagementResponse, inputTokens []string) (errTokens []string) {
	for _, info := range r.Errors {
		errTokens = append(errTokens, inputTokens[info.Index])
	}
	return
}

// IsAllowedTopic ...
func IsAllowedTopic(topic string) bool {
	return funk.ContainsString(topics, topic)
}


// UnsubscribeTokenFromTopic ...
func UnsubscribeTokenFromTopic(topic string, tokens []string) error {
	tokens = removeEmptyStrings(tokens)
	if !IsAllowedTopic(topic) {
		return errors.New(errTopicNotAllowed)
	}
	if len(tokens) <= 0 {
		return errors.New(errEmptyTokens)
	}
	res, err := clientMessage.UnsubscribeFromTopic(context.Background(), tokens, topic)
	if err != nil {
		fmt.Printf("Unsubscribe to topic %s error: %v", topic, err)
		return err
	}
	fmt.Printf("Unsubscribe tokens to topic %s: success %d, failed %d \n", topic, res.SuccessCount, res.FailureCount)
	if len(res.Errors) > 0 {
		return errors.New(res.Errors[0].Reason)
	}
	return nil
}