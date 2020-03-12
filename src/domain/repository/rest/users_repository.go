package rest

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/r-zareba/bookstore_oauth-api/src/domain/model/users"
	"github.com/r-zareba/bookstore_oauth-api/src/utils/errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type UsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

func NewUsersRepository() UsersRepository {
	return &usersRepository{}
}

type usersRepository struct{}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("Invalid Rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.InternalServerError("Invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	err := json.Unmarshal(response.Bytes(), &user)
	if err != nil {
		return nil, errors.InternalServerError("Error when trying to unmarshal users response")
	}
	return &user, nil
}
