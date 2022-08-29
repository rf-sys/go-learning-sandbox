package runtime

import (
	"awesomeProject1/config"
	"awesomeProject1/db"
	"awesomeProject1/repository"
	"github.com/jmoiron/sqlx"
)

var Config config.Config
var Database *sqlx.DB
var UserRepository repository.UserRepository

func InitRuntimeEnvironment() error {
	var err error

	Config, err = config.NewConfig()
	if err != nil {
		return err
	}

	Database, err = db.NewDb(Config)
	if err != nil {
		return err
	}

	UserRepository = repository.NewUserRepository(Database)

	return nil
}
