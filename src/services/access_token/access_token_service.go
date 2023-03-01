package service

import (
	domain "github.com/luizmoitinho/bookstore_oauth_api/src/domain/domain"
	"github.com/luizmoitinho/bookstore_oauth_api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*domain.AcessToken, *errors.RestError)
	Create(domain.AcessToken) *errors.RestError
	UpdateExpirationTime(domain.AcessToken) *errors.RestError
}

type AccessToken interface {
	GetByID(string) (*domain.AcessToken, *errors.RestError)
	Create(domain.AcessToken) *errors.RestError
	UpdateExpirationTime(domain.AcessToken) *errors.RestError
}

type service struct {
	repository Repository
}

func NewAccessToken(repositoryInjection Repository) AccessToken {
	return &service{
		repository: repositoryInjection,
	}
}

func (s *service) GetByID(accessTokenID string) (*domain.AcessToken, *errors.RestError) {
	at := domain.AcessToken{Token: accessTokenID}
	if err := at.IsTokenValid(); err != nil {
		return nil, err
	}

	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at domain.AcessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at domain.AcessToken) *errors.RestError {

	if err := at.IsTokenValid(); err != nil {
		return err
	}
	if err := at.IsExpiresValid(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
