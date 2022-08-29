package repository

import (
	"awesomeProject1/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(database *sqlx.DB) UserRepository {
	return UserRepository{
		db: database,
	}
}

func (repository UserRepository) FindAll() ([]model.User, error) {
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

func (repository UserRepository) Create(user model.User) error {
	log.Debug().Msgf("creating password hash for user %s", user.Username)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return err
	}

	var newUser model.User
	newUser.Username = user.Username
	newUser.Password = string(hash)
	newUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	log.Debug().Msgf("inserting new user %s", user.Username)
	_, err = repository.db.NamedExec("INSERT INTO users (username, password, created_at) VALUES (:username, :password, :created_at) RETURNING id", newUser)
	if err != nil {
		return err
	}

	return nil
}
