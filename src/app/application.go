package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luizmoitinho/bookstore_oauth_api/src/clients/cassandra"
	"github.com/luizmoitinho/bookstore_oauth_api/src/http"
	"github.com/luizmoitinho/bookstore_oauth_api/src/repository/db"
	rest "github.com/luizmoitinho/bookstore_oauth_api/src/repository/rest"
	service "github.com/luizmoitinho/bookstore_oauth_api/src/services/access_token"
	"gopkg.in/resty.v1"
)

var (
	router = gin.Default()
)

func StartApplication() {
	cassandra.InitCluster()
	session := cassandra.GetSession()
	defer session.Close()

	client := rest.NewClient(resty.New())
	accessTokenHandler := http.NewHandler(service.NewAccessToken(rest.NewRepository(client), db.New()))

	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetByID)
	router.POST("/oauth/access_token/", accessTokenHandler.Create)
	router.PUT("/oauth/access_token/:access_token_id", accessTokenHandler.UpdateExpirationTime)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
