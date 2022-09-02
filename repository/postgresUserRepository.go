package repository

import (
	"awesomeProject1/log"
	"awesomeProject1/model"
	"github.com/jmoiron/sqlx"
)

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
	sql := "SELECT * FROM users"

	// select all users from the database and put them into "users" array
	var users []model.User
	err := repository.db.Select(&users, sql)
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []model.User{}
	}

	return users, nil
}

func (repository PostgresUserRepository) FindOne(id int) (model.User, error) {
	sql := `SELECT id, username, password, created_at FROM users WHERE id = $1`

	var user model.User
	err := repository.db.Get(&user, sql, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repository PostgresUserRepository) Insert(user model.User) (int, error) {
	sql := "INSERT INTO users (username, password, created_at) VALUES (:username, :password, :created_at) RETURNING id"

	// prepare SQL query to insert new user
	var id int
	stmt, err := repository.db.PrepareNamed(sql)
	if err != nil {
		return id, err
	}

	// execute SQL query to insert new user
	err = stmt.Get(&id, user)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (repository PostgresUserRepository) Update(user model.User) error {
	sql := "UPDATE users SET (username, password) = (:username, :password) WHERE id = :id"

	_, err := repository.db.NamedExec(sql, user)
	if err != nil {
		return err
	}

	return nil
}
