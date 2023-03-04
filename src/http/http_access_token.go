package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	service "github.com/luizmoitinho/bookstore_oauth_api/src/services/access_token"
	"github.com/luizmoitinho/bookstore_utils/rest_errors"
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
	service service.AccessToken
}

func NewHandler(service service.AccessToken) AccessTokenHandler {
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
	var request domain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequestError := rest_errors.NewBadRequestError(INVALID_JSON_BODY)
		c.JSON(badRequestError.Status, badRequestError)
		return
	}

	accessToken, errCreate := handler.service.Create(request)
	if errCreate != nil {
		c.JSON(errCreate.Status, errCreate)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	accesTokenID := c.Param("access_token_id")
	atStored, err := handler.service.GetByID(accesTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var updateAt domain.AcessToken
	if err := c.ShouldBindJSON(&updateAt); err != nil {
		badRequestError := rest_errors.NewBadRequestError(INVALID_JSON_BODY)
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
