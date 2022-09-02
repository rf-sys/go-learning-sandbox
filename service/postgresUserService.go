package service

import (
	"awesomeProject1/model"
	"awesomeProject1/repository"
	"github.com/lib/pq"
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
	// create bcrypt hash from given password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		Username:  user.Username,
		Password:  string(hash),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// store used in the database
	id, err := service.repository.Insert(newUser)
	if err != nil {
		// see https://www.postgresql.org/docs/14/errcodes-appendix.html
		if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" && e.Constraint == "unique_username_constraint" {
			return model.User{}, ErrUserAlreadyExists
		}

		return model.User{}, err
	}

	if err != nil {
		return model.User{}, err
	}

	newUser.ID = id

	return newUser, nil
}

func NewPostgresUserService(repository repository.UserRepository) PostgresUserService {
	return PostgresUserService{repository: repository}
}
