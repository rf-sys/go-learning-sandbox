package runtime

import (
	"awesomeProject1/config"
	"awesomeProject1/db"
	"awesomeProject1/log"
	"awesomeProject1/repository"
	"github.com/jmoiron/sqlx"
)

var Config config.Config
var Logger log.Logger
var Database *sqlx.DB
var UserRepository repository.UserRepository

func InitRuntimeEnvironment() error {
	var err error

	Config, err = config.NewConfig()
	if err != nil {
		return err
	}

	Logger = log.NewZeroLogger()

	Database, err = db.NewDb(Config)
	if err != nil {
		return err
	}

	UserRepository = repository.NewUserRepository(Database, Logger)

	return nil
}
