package service

import (
	"github.com/r-zareba/bookstore_oauth-api/src/domain/model"
	"github.com/r-zareba/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetTokenById(string) (*model.AccessToken, *errors.RestError)
}

type AccessTokenService interface {
	GetTokenById(string) (*model.AccessToken, *errors.RestError)
}

func NewAccessTokenService(repository Repository) AccessTokenService {
	return &accessTokenService{repository}
}

type accessTokenService struct {
	repository Repository
}

func (service *accessTokenService) GetTokenById(tokenId string) (*model.AccessToken, *errors.RestError) {
	tokenId = strings.TrimSpace(tokenId)
	if len(tokenId) == 0 {
		return nil, errors.BadRequestError("Invalid Access Token ID")
	}

	accessToken, err := service.repository.GetTokenById(tokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}


