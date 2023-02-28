package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luizmoitinho/bookstore_oauth_api/src/clients/cassandra"
	"github.com/luizmoitinho/bookstore_oauth_api/src/domain/access_token"
	"github.com/luizmoitinho/bookstore_oauth_api/src/http"
	"github.com/luizmoitinho/bookstore_oauth_api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	accessTokenHandler := http.NewHandler(access_token.NewService(db.New()))

	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetByID)
	router.POST("/oauth/access_token/", accessTokenHandler.Create)
	router.PUT("/oauth/access_token/:access_token_id", accessTokenHandler.UpdateExpirationTime)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
