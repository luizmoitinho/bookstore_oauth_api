package access_token

import "time"

const (
	EXPIRATION_TIME = 24 // 24 hours
)

type AcessToken struct {
	Token    string `json:"access_token"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"`
	Expires  int64  `json:"expires"`
}

func GetNewAccesToken() *AcessToken {
	return &AcessToken{
		Expires: time.Now().UTC().Add(EXPIRATION_TIME + time.Hour).Unix(),
	}
}

func (at AcessToken) IsExpired() bool {
	now := time.Now().UTC()
	expireationTime := time.Unix(at.Expires, 0)

	return expireationTime.Before(now)
}
