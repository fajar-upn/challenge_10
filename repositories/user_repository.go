package repositories

import (
	"challenge_10/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(entity.User) (*entity.User, error)
	ValidateUser(string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user entity.User) (*entity.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) ValidateUser(username string) (*entity.User, error) {

	var user entity.User

	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
