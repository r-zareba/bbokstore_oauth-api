package application

import (
	"github.com/gin-gonic/gin"
	"github.com/r-zareba/bookstore_oauth-api/src/domain/repository/cassandra_db"
	"github.com/r-zareba/bookstore_oauth-api/src/domain/service"
	"github.com/r-zareba/bookstore_oauth-api/src/interface/http_handler"
)

var (
	router = gin.Default()
)

func StartApplication() {
	accessTokenRepository := cassandra_db.NewCassandraRepository()
	accessTokenService := service.NewAccessTokenService(accessTokenRepository)
	accessTokenHandler := http_handler.NewAccessTokenHttpHandler(accessTokenService)

	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetTokenById)

	router.Run(":8080")

}