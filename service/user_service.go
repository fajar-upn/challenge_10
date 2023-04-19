package service

import (
	"challenge_10/entity"
	"challenge_10/repositories"
	"errors"
	"time"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateClientUser(entity.PayloadUser) (*entity.User, error)
	CreateAdminUser(entity.PayloadUser) (*entity.User, error)
	ValidateUser(username, password string) (*entity.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (us *userService) CreateClientUser(payload entity.PayloadUser) (*entity.User, error) {
	uuid, _ := uuid.NewV4()

	pass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	userEntity := entity.User{
		ID:           uuid,
		Username:     payload.Username,
		Password:     string(pass),
		Access_level: "client",
		Created_at:   time.Now().Local(),
		Updated_at:   time.Now().Local(),
	}

	clientUser, err := us.repository.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}

	return clientUser, nil
}

func (us *userService) CreateAdminUser(payload entity.PayloadUser) (*entity.User, error) {
	uuid, _ := uuid.NewV4()

	pass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	stringPass := string(pass)

	userEntity := entity.User{
		ID:           uuid,
		Username:     payload.Username,
		Password:     stringPass,
		Access_level: "admin",
		Created_at:   time.Now().Local(),
		Updated_at:   time.Now().Local(),
	}

	adminUser, err := us.repository.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

func (us *userService) ValidateUser(username, password string) (*entity.User, error) {

	getUserByUsername, err := us.repository.ValidateUser(username)
	if err != nil {
		return nil, errors.New("Wrong Username or Password!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(getUserByUsername.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Wrong Username or Password!")
	}

	return getUserByUsername, nil
}
