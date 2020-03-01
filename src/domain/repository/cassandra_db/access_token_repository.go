package cassandra_db

import (
	"github.com/r-zareba/bookstore_oauth-api/src/domain/model"
	"github.com/r-zareba/bookstore_oauth-api/src/utils/errors"
)

type CassandraRepository interface {
	GetTokenById(string) (*model.AccessToken, *errors.RestError)
}

func NewCassandraRepository() CassandraRepository {
	return &cassandraRepository{}
}

type cassandraRepository struct {
}

func (repository *cassandraRepository) GetTokenById(id string) (*model.AccessToken, *errors.RestError) {
	return nil, errors.InternalServerError("Method not implemented!")
}
