package service

import (
	"github.com/r-zareba/bookstore_oauth-api/src/domain/model/access_token"
	"github.com/r-zareba/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetTokenById(string) (*access_token.AccessToken, *errors.RestError)
	CreateToken(access_token.AccessToken) *errors.RestError
	UpdateExpiresIn(token access_token.AccessToken) *errors.RestError
}

type AccessTokenService interface {
	GetTokenById(string) (*access_token.AccessToken, *errors.RestError)
	CreateToken(access_token.AccessToken) *errors.RestError
	UpdateExpiresIn(access_token.AccessToken) *errors.RestError
}

func NewAccessTokenService(repository Repository) AccessTokenService {
	return &accessTokenService{repository}
}

type accessTokenService struct {
	repository Repository
}

func (s *accessTokenService) GetTokenById(tokenId string) (*access_token.AccessToken, *errors.RestError) {
	tokenId = strings.TrimSpace(tokenId)
	if len(tokenId) == 0 {
		return nil, errors.BadRequestError("Invalid Access Token ID")
	}

	accessToken, err := s.repository.GetTokenById(tokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *accessTokenService) CreateToken(token access_token.AccessToken) *errors.RestError {
	valErr := token.Validate()
	if valErr != nil {
		return valErr
	}
	return s.repository.CreateToken(token)
}

func (s *accessTokenService) UpdateExpiresIn(token access_token.AccessToken) *errors.RestError {
	valErr := token.Validate()
	if valErr != nil {
		return valErr
	}
	return s.repository.UpdateExpiresIn(token)
}


