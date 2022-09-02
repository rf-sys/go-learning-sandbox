package service

import (
	"awesomeProject1/model"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	Create(user model.User) (model.User, error)
}
