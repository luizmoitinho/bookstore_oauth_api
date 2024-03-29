package domain

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccesTokenConstants(t *testing.T) {
	assert.EqualValues(t, EXPIRATION_TIME, 24, "expiration time should be 24 hours")
}

func TestGetNewAccesToken(t *testing.T) {
	//assert
	expectedExpires := time.Now().UTC().Add(24 + time.Hour).Unix()
	//act
	at := NewAccessToken(1)

	//assert
	assert.NotNil(t, at, "new acess token was returned a nil pointer")
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.Equal(t, "", at.Token, "new access token should not have a defined access token id")
	assert.True(t, at.UserID == 1, "new access token should must be return a user id equals to 1")
	assert.EqualValues(t, expectedExpires, at.Expires, fmt.Sprintf("new access token return %v and was expected %v", at.Expires, expectedExpires))
}

func TestIsExpired(t *testing.T) {
	//act
	at := AcessToken{}

	//assert
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 + time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should not be expired")
}
