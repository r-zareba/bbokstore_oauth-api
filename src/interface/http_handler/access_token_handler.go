package http_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/r-zareba/bookstore_oauth-api/src/domain/service"
	"net/http"
)

type AccessTokenHandler interface {
	GetTokenById(ctx *gin.Context)
}

func NewAccessTokenHttpHandler(service service.AccessTokenService) AccessTokenHandler {
	return &accessTokenHttpHandler{service}
}

type accessTokenHttpHandler struct {
	service service.AccessTokenService
}

func (handler *accessTokenHttpHandler) GetTokenById(ctx *gin.Context) {
	accessTokenId := ctx.Param("access_token_id")

	accessToken, err := handler.service.GetTokenById(accessTokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}





