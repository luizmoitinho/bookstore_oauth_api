package db

import (
	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	"github.com/luizmoitinho/bookstore_oauth_api/src/utils/errors"
)

func New() DatabaseRepository {
	return &dbRespository{}
}

type DatabaseRepository interface {
	GetByID(string) (*access_token.AcessToken, *errors.RestError)
}

type dbRespository struct {
}

func (db *dbRespository) GetByID(string) (*access_token.AcessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("database connection not implemented yet.")
}
