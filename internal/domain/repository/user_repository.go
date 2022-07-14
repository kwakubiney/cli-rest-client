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

// create User
func (u *UserRepository) CreateUser(user model.User) error {
	return u.db.Create(&user).Error
}

// Update user by username
func (u *UserRepository) UpdateUserByUsername(username string, user model.User) error {
	return u.db.Model(model.User{}).Where("username = ?", username).Updates(&user).Error
}
