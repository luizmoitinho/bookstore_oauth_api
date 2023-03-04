package rest_users

import (
	"encoding/json"
	"net/http"

	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/users"
	"github.com/luizmoitinho/bookstore_utils/rest_errors"
)

type RestUsersRepository interface {
	Login(string, string) (*users.User, *rest_errors.RestError)
}

type userRepository struct {
	client Client
}

func NewRepository(clientInjection Client) RestUsersRepository {
	return &userRepository{
		client: clientInjection,
	}
}

func (r *userRepository) Login(email, password string) (*users.User, *rest_errors.RestError) {
	login := users.Login{
		Email:    email,
		Password: password,
	}

	result, err := r.client.OAuthLoginRequest(&login)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to login user", err)
	}

	if result.StatusCode() >= http.StatusMultipleChoices {
		var restErr rest_errors.RestError
		err := json.Unmarshal(result.Body(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(result.Body(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unsmarshall users login response", err)
	}

	return &user, nil
}
