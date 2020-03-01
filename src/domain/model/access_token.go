package model

import "time"

const (
	expirationTime = 24 // hours
)

type AccessToken struct {
	Token     string `json:"access_token"`
	UserId    int64  `json:"user_id"`
	ClientId  int64  `json:"user_id"`
	ExpiresIn int64  `json:"user_id"`
}

func GetAccessToken() *AccessToken {
	return &AccessToken{
		Token:     "",
		UserId:    0,
		ClientId:  0,
		ExpiresIn: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(token.ExpiresIn, 0)
	return now.After(expirationTime)
}
