package service

import (
	domain "github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/users"
	"github.com/luizmoitinho/bookstore_utils/rest_errors"
)

type Repository interface {
	GetByID(string) (*domain.AcessToken, *rest_errors.RestError)
	Create(domain.AcessToken) *rest_errors.RestError
	UpdateExpirationTime(domain.AcessToken) *rest_errors.RestError
}

type RestUsersRepository interface {
	Login(string, string) (*users.User, *rest_errors.RestError)
}

type AccessToken interface {
	GetByID(string) (*domain.AcessToken, *rest_errors.RestError)
	Create(in domain.AccessTokenRequest) (*domain.AcessToken, *rest_errors.RestError)
	UpdateExpirationTime(domain.AcessToken) *rest_errors.RestError
}

type service struct {
	database Repository
	users    RestUsersRepository
}

func NewAccessToken(u RestUsersRepository, db Repository) AccessToken {
	return &service{
		database: db,
		users:    u,
	}
}

func (s *service) GetByID(accessTokenID string) (*domain.AcessToken, *rest_errors.RestError) {
	at := domain.AcessToken{Token: accessTokenID}
	if err := at.IsTokenValid(); err != nil {
		return nil, err
	}

	accessToken, err := s.database.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(in domain.AccessTokenRequest) (*domain.AcessToken, *rest_errors.RestError) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	user, err := s.users.Login(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	at := domain.NewAccessToken(user.Id)
	at.GenerateCrypto()

	if err := s.database.Create(*at); err != nil {
		return nil, err
	}
	return at, nil
}

func (s *service) UpdateExpirationTime(at domain.AcessToken) *rest_errors.RestError {

	if err := at.IsTokenValid(); err != nil {
		return err
	}
	if err := at.IsExpiresValid(); err != nil {
		return err
	}
	return s.database.UpdateExpirationTime(at)
}
