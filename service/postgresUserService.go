package service

import (
	"awesomeProject1/model"
	"awesomeProject1/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type PostgresUserService struct {
	repository repository.UserRepository
}

func (service PostgresUserService) GetAllUsers() ([]model.User, error) {
	return service.repository.FindAll()
}

func (service PostgresUserService) Create(user model.User) (model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		Username:  user.Username,
		Password:  string(hash),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	id, err := service.repository.Insert(newUser)
	if err != nil {
		return model.User{}, err
	}

	newUser.ID = id

	return newUser, nil
}

func NewPostgresUserService(repository repository.UserRepository) PostgresUserService {
	return PostgresUserService{repository: repository}
}
