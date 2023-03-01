package domain

import (
	"strings"
	"time"

	"github.com/luizmoitinho/bookstore_oauth_api/src/utils/errors"
)

const (
	EXPIRATION_TIME   = 24 // 24 hours
	INVALID_TOKEN_ID  = "invalid access token id"
	INVALID_USER_ID   = "invalid user id"
	INVALID_CLIENT_ID = "invalid client id"
	INVALID_EXPIRES   = "invalid expires"
)

type AcessToken struct {
	Token    string `json:"access_token"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
	Expires  int64  `json:"expires"`
}

func NewAccessToken() *AcessToken {
	return &AcessToken{
		Expires: time.Now().UTC().Add(EXPIRATION_TIME + time.Hour).Unix(),
	}

}

func (at *AcessToken) IsTokenValid() *errors.RestError {
	if len(strings.TrimSpace(at.Token)) == 0 {
		return errors.NewBadRequestError(INVALID_TOKEN_ID)
	}
	return nil
}

func (at *AcessToken) IsUserIdValid() *errors.RestError {
	if at.UserID <= 0 {
		return errors.NewBadRequestError(INVALID_USER_ID)
	}
	return nil
}

func (at *AcessToken) IsClientIdValid() *errors.RestError {
	if at.ClientID <= 0 {
		return errors.NewBadRequestError(INVALID_CLIENT_ID)
	}
	return nil
}

func (at *AcessToken) IsExpiresValid() *errors.RestError {
	if at.Expires <= 0 {
		return errors.NewBadRequestError(INVALID_EXPIRES)
	}
	return nil
}

func (at *AcessToken) Validate() *errors.RestError {
	if err := at.IsTokenValid(); err != nil {
		return err
	}
	if err := at.IsUserIdValid(); err != nil {
		return err
	}
	if err := at.IsClientIdValid(); err != nil {
		return err
	}
	if err := at.IsExpiresValid(); err != nil {
		return err
	}
	return nil
}

func (at AcessToken) IsExpired() bool {
	now := time.Now().UTC()
	expireationTime := time.Unix(at.Expires, 0)

	return expireationTime.Before(now)
}
