package access_token

import (
	"github.com/luizmoitinho/bookstore_oauth_api/src/utils/errors"
)

const (
	INVALID_TOKEN_ID = "invalid access token id"
)

type Repository interface {
	GetByID(string) (*AcessToken, *errors.RestError)
}

type Service interface {
	GetByID(string) (*AcessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

func NewService(repositoryInjection Repository) Service {
	return &service{
		repository: repositoryInjection,
	}
}

func (s *service) GetByID(accessTokenID string) (*AcessToken, *errors.RestError) {
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError(INVALID_TOKEN_ID)
	}
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
