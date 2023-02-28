package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	"github.com/luizmoitinho/bookstore_oauth_api/src/utils/errors"
)

const (
	INVALID_JSON_BODY = "invalid json body"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accesTokenID := c.Param("access_token_id")
	accessToken, err := handler.service.GetByID(accesTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AcessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		badRequestError := errors.NewBadRequestError(INVALID_JSON_BODY)
		c.JSON(badRequestError.Status, badRequestError)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}

func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	accesTokenID := c.Param("access_token_id")
	atStored, err := handler.service.GetByID(accesTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var updateAt access_token.AcessToken
	if err := c.ShouldBindJSON(&updateAt); err != nil {
		badRequestError := errors.NewBadRequestError(INVALID_JSON_BODY)
		c.JSON(badRequestError.Status, badRequestError)
		return
	}
	atStored.Expires = updateAt.Expires

	if err := handler.service.UpdateExpirationTime(*atStored); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, atStored)
}
