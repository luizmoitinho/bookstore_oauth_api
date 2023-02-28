package rest_users

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/users"
	"gopkg.in/resty.v1"
)

type Client interface {
	OAuthLoginRequest(login *users.Login) (*resty.Response, error)
}
type client struct {
	_http *resty.Client
}

func NewClient(c *resty.Client) Client {
	return &client{
		_http: c,
	}
}

func (c *client) OAuthLoginRequest(login *users.Login) (*resty.Response, error) {
	var timeoutDuration time.Duration
	usersTimeout, err := strconv.ParseInt(os.Getenv("REST_USERS_TIMEOUT"), 10, 64)
	if err != nil {
		usersTimeout = 100
	}
	timeoutDuration = time.Duration(usersTimeout) * time.Millisecond

	c._http.SetTimeout(timeoutDuration)

	resp, err := c._http.R().
		SetBody(login).
		Post(fmt.Sprintf("%s%s", os.Getenv("REST_USERS_BASE_URL"), os.Getenv("REST_USERS_URI")))

	return resp, err
}
