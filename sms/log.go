package sms

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	otpmodel "github.com/cashbagvn/module-otp/model"
)

// SaveMultiLog
func SaveMultiLog(ctx context.Context, col *mongo.Collection, smsLogs []interface{}) error {
	_, err := col.InsertMany(ctx, smsLogs)
	return err
}

// NewSMSLogRaw ...
func NewSMSLogRaw(requestBody, service, smsType, recipient, content, result, deviceId, source, purpose string, success bool) *otpmodel.SMSLogRaw {
	return &otpmodel.SMSLogRaw{
		ID:          primitive.NewObjectID(),
		RequestBody: requestBody,
		Service:     service,
		Type:        smsType,
		Recipient:   recipient,
		Content:     content,
		DeviceID:    deviceId,
		Source:      source,
		Purpose:     purpose,
		CreatedAt:   time.Now(),
		Success:     success,
		Result:      result,
	}
}
