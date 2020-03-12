package access_token

import (
	"github.com/r-zareba/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24 // hours
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// For password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// For client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AccessToken struct {
	Token     string `json:"token"`
	UserId    int64  `json:"user_id"`
	ClientId  int64  `json:"client_id"`
	ExpiresIn int64  `json:"expires_in"`
}

func (at *AccessTokenRequest) Validate() *errors.RestError {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.BadRequestError("Invalid grant type parameter")
	}

	return nil

}

func GetAccessToken() *AccessToken {
	return &AccessToken{
		Token:     "",
		UserId:    0,
		ClientId:  0,
		ExpiresIn: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (t *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(t.ExpiresIn, 0)
	return now.After(expirationTime)
}

func (t *AccessToken) Validate() *errors.RestError {
	t.Token = strings.TrimSpace(t.Token)
	if t.Token == "" {
		return errors.BadRequestError("Invalid token id")
	}
	if t.UserId <= 0 {
		return errors.BadRequestError("Invalid user id")
	}
	if t.ClientId <= 0 {
		return errors.BadRequestError("Invalid client id")
	}
	if t.ExpiresIn <= 0 {
		return errors.BadRequestError("Invalid token expiration time")
	}
	return nil
}
