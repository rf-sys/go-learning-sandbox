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

func (service PostgresUserService) Edit(payload model.User) error {
	// check if payload exists in the database before trying to edit it
	dbUser, err := service.GetOne(payload.ID)
	if err != nil {
		return ErrUserNotFound
	}

	user := model.User{}

	if payload.Username != "" && dbUser.Username != payload.Username {
		user.Username = payload.Username
	} else {
		user.Username = dbUser.Username
	}

	if payload.Password != "" && dbUser.Password != payload.Password {
		hash, err2 := createPasswordHash(payload.Password)
		if err2 != nil {
			return err
		}
		user.Password = hash
	} else {
		user.Password = dbUser.Password
	}

	user.ID = dbUser.ID

	err = service.repository.Update(user)
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
