package repository

import (
		"github.com/kwakubiney/canonical-take-home/internal/domain/model"
)

type UserRepository struct {
	baseUrl string
}

const Endpoint = "user"

func NewUserRepository(baseUrl string) *UserRepository {
	return &UserRepository{
		baseUrl,
	}
}

// fetch by keys/key
func (u *UserRepository) FindByKeys(keys []string) *model.User {
	// build URL
	

	// call API util to make web request
	
	// unmarshall response
	return &model.User{}
}

// fetch many
func (u *UserRepository) Find(numResults int)*model.User {
	// build URL

	// call api util

	// unmarshal results into slice
	
	return &model.User{}
}
