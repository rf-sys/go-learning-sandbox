package repository

import (
	"awesomeProject1/model"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Insert(user model.User) (int, error)
}
