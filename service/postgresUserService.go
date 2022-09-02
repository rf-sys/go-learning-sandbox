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

func (service PostgresUserService) GetOne(id int) (model.User, error) {
	return service.repository.FindOne(id)
}

func (service PostgresUserService) Create(user model.User) (model.User, error) {
	// create bcrypt hash from given password
	hash, err := createPasswordHash(user.Password)
	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		Username:  user.Username,
		Password:  hash,
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

	newUser.ID = id

	return newUser, nil
}

func (service PostgresUserService) Edit(id int, user model.User) error {
	// check if user exists in the database before trying to edit it
	dbUser, err := service.GetOne(id)
	if err != nil {
		return ErrUserNotFound
	}

	updatedUser := model.User{}

	if user.Username != "" && dbUser.Username != user.Username {
		updatedUser.Username = user.Username
	} else {
		updatedUser.Username = dbUser.Username
	}

	if user.Password != "" && dbUser.Password != user.Password {
		hash, e := createPasswordHash(user.Password)
		if e != nil {
			return err
		}
		updatedUser.Password = hash
	} else {
		updatedUser.Password = dbUser.Password
	}

	updatedUser.ID = id
	err = service.repository.Update(updatedUser)
	if err != nil {
		return err
	}

	return nil
}

func createPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func NewPostgresUserService(repository repository.UserRepository) PostgresUserService {
	return PostgresUserService{repository: repository}
}
