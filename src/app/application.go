package app

import (
	"github.com/gin-gonic/gin"
	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	"github.com/luizmoitinho/bookstore_oauth_api/src/http"
	"github.com/luizmoitinho/bookstore_oauth_api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	accessTokenHandler := http.NewHandler(access_token.NewService(db.New()))

	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetByID)
	router.Run(":8080")
}
