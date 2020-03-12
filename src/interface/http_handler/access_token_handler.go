package http_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/r-zareba/bookstore_oauth-api/src/domain/model/access_token"
	"github.com/r-zareba/bookstore_oauth-api/src/domain/service"
	"github.com/r-zareba/bookstore_oauth-api/src/utils/errors"
	"net/http"
)

type AccessTokenHandler interface {
	GetTokenById(*gin.Context)
	CreateToken(*gin.Context)
	UpdateExpiresIn(*gin.Context)
}

func NewAccessTokenHttpHandler(service service.AccessTokenService) AccessTokenHandler {
	return &accessTokenHttpHandler{service}
}

type accessTokenHttpHandler struct {
	service service.AccessTokenService
}

func (h *accessTokenHttpHandler) GetTokenById(ctx *gin.Context) {
	accessTokenId := ctx.Param("access_token_id")
	accessToken, err := h.service.GetTokenById(accessTokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHttpHandler) CreateToken(ctx *gin.Context) {
	var accessToken access_token.AccessToken

	err := ctx.ShouldBindJSON(&accessToken)
	if err != nil {
		restErr := errors.BadRequestError("Invalid JSON body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	createErr := h.service.CreateToken(accessToken)
	if createErr != nil {
		ctx.JSON(createErr.Status, createErr)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHttpHandler) UpdateExpiresIn(ctx *gin.Context) {

}




