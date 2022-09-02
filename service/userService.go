package service

import (
	"awesomeProject1/model"
	"errors"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	Create(user model.User) (model.User, error)
}

var ErrUserAlreadyExists = errors.New("user already exists")
