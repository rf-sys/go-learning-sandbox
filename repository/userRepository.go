package repository

import (
	"awesomeProject1/model"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindOne(id int) (model.User, error)
	Insert(user model.User) (int, error)
	Update(user model.User) error
}
