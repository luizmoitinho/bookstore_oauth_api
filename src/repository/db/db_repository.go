package db

import (
	"github.com/gocql/gocql"
	"github.com/luizmoitinho/bookstore_oauth_api/src/clients/cassandra"
	domain "github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	"github.com/luizmoitinho/bookstore_utils/rest_errors"
)

const (
	QUERY_SELECT_BY_ACCESS_TOKEN = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	QUERY_INSERT_ACCESS_TOKEN    = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	QUERY_UPDATE_EXPIRE          = "UPDATE access_tokens SET expires = ? WHERE access_token = ?"
)

func New() DatabaseRepository {
	return &dbRespository{}
}

type DatabaseRepository interface {
	GetByID(string) (*domain.AcessToken, *rest_errors.RestError)
	Create(domain.AcessToken) *rest_errors.RestError
	UpdateExpirationTime(domain.AcessToken) *rest_errors.RestError
}

type dbRespository struct {
}

func (db *dbRespository) UpdateExpirationTime(at domain.AcessToken) *rest_errors.RestError {
	if err := cassandra.GetSession().Query(QUERY_UPDATE_EXPIRE, at.Expires, at.Token).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when trying update expiration time at database", err)
	}
	return nil
}

func (db *dbRespository) Create(at domain.AcessToken) *rest_errors.RestError {
	if err := cassandra.GetSession().Query(QUERY_INSERT_ACCESS_TOKEN, at.Token, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when trying to save access token in database", err)
	}
	return nil
}

func (db *dbRespository) GetByID(id string) (*domain.AcessToken, *rest_errors.RestError) {
	var result = domain.AcessToken{}
	if err := cassandra.GetSession().Query(QUERY_SELECT_BY_ACCESS_TOKEN, id).Scan(&result.Token, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error when trying get token by id at database", err)
	}
	return &result, nil
}
