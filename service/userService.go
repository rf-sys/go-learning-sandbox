package service

import (
	"awesomeProject1/log"
	"awesomeProject1/model"
	"awesomeProject1/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	FindAll() ([]model.User, error)
	Create(user model.User) (model.User, error)
}

type PostgresUserService struct {
	repository repository.UserRepository
	logger     log.Logger
}

func (service PostgresUserService) FindAll() ([]model.User, error) {
	return service.repository.FindAll()
}

func (service PostgresUserService) Create(user model.User) (model.User, error) {
	service.logger.Debug(fmt.Sprintf("creating password hash for user %s", user.Username))
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return model.User{}, err
	}

	var newUser model.User
	newUser.Username = user.Username
	newUser.Password = string(hash)
	newUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	service.logger.Debug(fmt.Sprintf("inserting new user %s", user.Username))
	return service.repository.Create(newUser)
}

func NewPostgresUserService(repository repository.UserRepository, logger log.Logger) UserService {
	return PostgresUserService{repository: repository, logger: logger}
}
