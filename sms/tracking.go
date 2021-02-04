package sms

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	otpconfig "github.com/cashbagvn/module-otp/config"
	otpmodel "github.com/cashbagvn/module-otp/model"
)

// SaveTracking
func SaveTracking(ctx context.Context, col *mongo.Collection, raw otpmodel.OtpRaw) error {
	_, err := col.InsertOne(ctx, raw)
	return err
}

// VerifyOtp
func VerifyOtp(ctx context.Context, col *mongo.Collection, req otpmodel.VerifyOTPRequest) (isValid bool) {
	if !isValidTracker(req.Tracker) {
		fmt.Println("Tracker: ", req.Tracker)
		return
	}
	total, _ := col.CountDocuments(ctx, bson.M{
		"expireAt": bson.M{"$gt": time.Now()},
		"code":     req.Code,
		"tracker":  req.Tracker,
		"to":       req.Receipt,
	})
	if total > 0 {
		cond := bson.M{
			"code":    req.Code,
			"tracker": req.Tracker,
			"to":      req.Receipt,
		}
		go col.DeleteMany(context.Background(), cond)
	}
	return total > 0
}

func isValidTracker(s string) bool {
	re := regexp.MustCompile(`(?m)[a-z]+\:[a-z\-]+\:[0-9a-fA-F]{24}`)
	return re.Match([]byte(s))
}

func isValidOTPType(otpType string) bool {
	validTypes := []string{
		otpconfig.OtpTypeSMS,
		otpconfig.OtpTypeEmail,
	}
	return funk.ContainsString(validTypes, otpType)
}

// GenerateOTPTracker
func GenerateOTPTracker(serviceName, collection string, targetID primitive.ObjectID) string {
	return fmt.Sprintf("%s:%s:%s", serviceName, collection, targetID.Hex())
}

// NewOtpBSON ...
func NewOtpBSON(to []string, otpType, tracker string) (otp otpmodel.OtpRaw) {
	now := time.Now()
	otp = otpmodel.OtpRaw{
		ID:        primitive.NewObjectID(),
		To:        to,
		ExpireAt:  now.Add(time.Minute * 30),
		Type:      otpType,
		CreatedAt: now,
		Tracker:   tracker,
	}
	rand.Seed(now.UnixNano())
	otp.Code = fmt.Sprintf("%d", rand.Intn(899999)+100000)
	return
}
