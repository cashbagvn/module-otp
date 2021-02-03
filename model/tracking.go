package otpmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OtpRaw
type OtpRaw struct {
	ID        primitive.ObjectID `bson:"_id"`
	To        []string           `bson:"to"`
	ExpireAt  time.Time          `bson:"expireAt"`
	Code      string             `bson:"code"`
	Type      string             `bson:"type"`
	CreatedAt time.Time          `bson:"createdAt"`
	Tracker   string             `bson:"tracker"`
}

// VerifyOTPRequest
type VerifyOTPRequest struct {
	Receipt string
	Tracker string
	Code    string
}
