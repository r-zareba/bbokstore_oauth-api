package rest

import (
	"encoding/json"
	"errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/r-zareba/bookstore_oauth-api/src/domain/model/users"
	"github.com/r-zareba/bookstore_utils-go/rest_errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type UsersRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestError)
}

func NewUsersRepository() UsersRepository {
	return &usersRepository{}
}

type usersRepository struct{}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.InternalServerError("Invalid Rest client response when trying to login user",
			errors.New("response error"))
	}

	if response.StatusCode > 299 {
		var restErr rest_errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.InternalServerError("Invalid error interface when trying to login user0",
				errors.New("response error"))
		}
		return nil, &restErr
	}

	var user users.User
	err := json.Unmarshal(response.Bytes(), &user)
	if err != nil {
		return nil, rest_errors.InternalServerError("Error when trying to unmarshal users response",
			errors.New("response error"))
	}
	return &user, nil
}
