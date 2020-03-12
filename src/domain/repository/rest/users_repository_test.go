package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

//Runs before every test case
func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserApiTimeout(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "test@gmail.com", "password": "test123"`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("test@gmail.com", "test123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid Rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "test@gmail.com", "password": "test123"`,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"message": "Invalid login", "status": "fake"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("test@gmail.com", "test123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "test@gmail.com", "password": "test123"`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Invalid login credentials", "status": 404, "error": "not_found"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("test@gmail.com", "test123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Invalid login credentials", err.Message)
}

func TestLoginInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "test@gmail.com", "password": "test123"`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "Fake ID", "first_name": "Jan", "last_name": "Kowalski", "email": "test@gmail.com"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("test@gmail.com", "test123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error when trying to unmarshal users response", err.Message)
}

func TestLoginNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email": "test@gmail.com", "password": "test123"`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "Jan", "last_name": "Kowalski", "email": "test@gmail.com"}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("test@gmail.com", "test123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 1)
	assert.EqualValues(t, user.FirstName, "Jan")
	assert.EqualValues(t, user.LastName, "Kowalski")
	assert.EqualValues(t, user.Email, "test@gmail.com")
}
