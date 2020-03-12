package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time for access token should be 24 hours")
}

func TestGetAccessToken(t *testing.T) {
	token := GetAccessToken()

	assert.EqualValues(t, token.Token, "", "New token should not have defined access token id")
	assert.EqualValues(t, token.UserId, 0, "New access token should not have associated user id")
	assert.False(t, token.IsExpired(), false, "New token should not be expired")
}

func TestAccessToken_IsExpired(t *testing.T) {
	token := AccessToken{}
	assert.True(t, token.IsExpired(), true, "Empty Token should be expired by default")
	token.ExpiresIn = time.Now().UTC().Add(1 * time.Hour).Unix()
	assert.False(t, token.IsExpired(), false, "Adding 1 hour of expiration time should made token be not expired")
}
