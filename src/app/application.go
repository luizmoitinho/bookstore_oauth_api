package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luizmoitinho/bookstore_oauth_api/src/clients/cassandra"
	"github.com/luizmoitinho/bookstore_oauth_api/src/http"
	"github.com/luizmoitinho/bookstore_oauth_api/src/repository/db"
	service "github.com/luizmoitinho/bookstore_oauth_api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	cassandra.InitCluster()
	session := cassandra.GetSession()
	defer session.Close()

	accessTokenHandler := http.NewHandler(service.AccessToken(db.New()))

	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetByID)
	router.POST("/oauth/access_token/", accessTokenHandler.Create)
	router.PUT("/oauth/access_token/:access_token_id", accessTokenHandler.UpdateExpirationTime)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
