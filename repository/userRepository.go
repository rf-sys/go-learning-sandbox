package repository

import (
	"awesomeProject1/log"
	"awesomeProject1/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Insert(user model.User) (int, error)
}

type PostgresUserRepository struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewPostgresUserRepository(database *sqlx.DB) PostgresUserRepository {
	return PostgresUserRepository{
		db: database,
	}
}

func (repository PostgresUserRepository) FindAll() ([]model.User, error) {
	var users []model.User

	err := repository.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []model.User{}
	}

	return users, nil
}

func (repository PostgresUserRepository) Insert(user model.User) (int, error) {
	var id int

	stmt, err := repository.db.PrepareNamed("INSERT INTO users (username, password, created_at) VALUES (:username, :password, :created_at) RETURNING id")
	if err != nil {
		return id, err
	}

	err = stmt.Get(&id, user)
	if err != nil {
		return id, err
	}

	return id, nil
}
