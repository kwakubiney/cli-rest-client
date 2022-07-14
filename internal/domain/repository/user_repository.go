package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

// fetch by keys/key
func (u *UserRepository) FindByKeys(keys []string, c *gin.Context) *model.User {
	// build URL

	// call API util to make web request

	// unmarshall response
	return &model.User{}
}

// fetch many
func (u *UserRepository) Find(numResults int) *model.User {
	// build URL

	// call api util

	// unmarshal results into slice

	return &model.User{}
}
