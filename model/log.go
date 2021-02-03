package otpmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SMSLogRaw ...
type SMSLogRaw struct {
	ID        primitive.ObjectID `bson:"_id"`
	Service   string             `bson:"service"` // vietguys, ...
	Carrier   string             `bson:"carrier"` // mobifone, viettel, vinaphone, ...
	Type      string             `bson:"type"`    // otp, ...
	Recipient string             `bson:"recipient"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"createdAt"`
	Success   bool               `bson:"success"`
	Result    string             `bson:"result"`
}

// SetCarrier ....
func (log *SMSLogRaw) SetCarrier(c string) *SMSLogRaw {
	log.Carrier = c
	return log
}
