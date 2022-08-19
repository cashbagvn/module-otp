package otpmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SMSLogRaw ...
type SMSLogRaw struct {
	ID          primitive.ObjectID `bson:"_id"`
	RequestBody string             `bson:"requestBody,omitempty"` // Optional
	Service     string             `bson:"service"`               // vietguys, ...
	Carrier     string             `bson:"carrier"`               // mobifone, viettel, vinaphone, ...
	Type        string             `bson:"type"`                  // otp, ...
	Recipient   string             `bson:"recipient"`
	Content     string             `bson:"content"`
	Purpose     string             `bson:"purpose"`
	Source      string             `bson:"source"`
	CreatedAt   time.Time          `bson:"createdAt"`
	Success     bool               `bson:"success"`
	DeviceID    string             `bson:"deviceId,omitempty"`
	Result      string             `bson:"result"`
}

// SetCarrier ....
func (log *SMSLogRaw) SetCarrier(c string) *SMSLogRaw {
	log.Carrier = c
	return log
}
