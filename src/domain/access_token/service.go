package access_token

import (
	"github.com/luizmoitinho/bookstore_oauth_api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AcessToken, *errors.RestError)
	Create(AcessToken) *errors.RestError
	UpdateExpirationTime(AcessToken) *errors.RestError
}

type Service interface {
	GetByID(string) (*AcessToken, *errors.RestError)
	Create(AcessToken) *errors.RestError
	UpdateExpirationTime(AcessToken) *errors.RestError
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
	at := AcessToken{Token: accessTokenID}
	if err := at.IsTokenValid(); err != nil {
		return nil, err
	}

	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AcessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AcessToken) *errors.RestError {
	if err := at.IsTokenValid(); err != nil {
		return err
	}
	if err := at.IsExpiresValid(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
