package repository

import (
	"fmt"
	"log"

	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
	"gorm.io/gorm"
)
type UserNotFoundError struct {
    msg string
}

func (e *UserNotFoundError) Error() string { return e.msg }
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
	db :=  u.db.Model(model.User{}).Where("username = ?", username).Find(&user)
	if db.RowsAffected == 0 {
		return  &UserNotFoundError{"no record found for username given"}
	}
	err := db.Updates(&user).Error
	if err != nil {
		log.Println(db.Error)
		return err
	}
	return nil
}

func (u *UserRepository) FilterUser(by string, where string) (*model.User, error) {
	var user model.User
	db := u.db.Where(fmt.Sprintf("%s = ?", by), where).Find(&user)
	if db.RowsAffected == 0 {
		return &user, &UserNotFoundError{"no record found for username given"}
	}
	if db.Error != nil {
		log.Println(db.Error)
		return nil, db.Error
	}
	return &user, db.Error
}
