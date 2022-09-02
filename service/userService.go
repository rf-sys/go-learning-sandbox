package service

import (
	"awesomeProject1/model"
	"errors"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetOne(id int) (model.User, error)
	Create(user model.User) (model.User, error)
	Edit(user model.User) error
}

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")
