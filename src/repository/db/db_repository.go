package db

import (
	"github.com/gocql/gocql"
	"github.com/luizmoitinho/bookstore_oauth_api/src/clients/cassandra"
	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	"github.com/luizmoitinho/bookstore_oauth_api/src/utils/errors"
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
	GetByID(string) (*access_token.AcessToken, *errors.RestError)
	Create(access_token.AcessToken) *errors.RestError
	UpdateExpirationTime(access_token.AcessToken) *errors.RestError
}

type dbRespository struct {
}

func (db *dbRespository) UpdateExpirationTime(at access_token.AcessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(QUERY_UPDATE_EXPIRE, at.Expires, at.Token).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (db *dbRespository) Create(at access_token.AcessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(QUERY_INSERT_ACCESS_TOKEN, at.Token, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (db *dbRespository) GetByID(id string) (*access_token.AcessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result = access_token.AcessToken{}
	if err := session.Query(QUERY_SELECT_BY_ACCESS_TOKEN, id).Scan(&result.Token, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
