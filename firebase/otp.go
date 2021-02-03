package firebase

import "firebase.google.com/go/auth"

// VerifyIDToken ...
func VerifyIDToken(idToken string) (*auth.Token, error) {
	return client.VerifyIDToken(ctx, idToken)
}

// GetUserInfoByUID ...
func GetUserInfoByUID(uid string) (*auth.UserRecord, error) {
	return client.GetUser(ctx, uid)
}
