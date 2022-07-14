package repository

import (
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
func (u *UserRepository) CreateUser(user model.User) error {
	return u.db.Create(&user).Error
}

// fetch many
func (u *UserRepository) Find(numResults int) *model.User {
	// build URL

	// call api util

	// unmarshal results into slice

	return &model.User{}
}
