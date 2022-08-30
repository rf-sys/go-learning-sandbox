package db

import (
	"awesomeProject1/config"
	"awesomeProject1/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewDb connects to the database and verifies connection with ping
func NewDb(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.Datasource)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.Database.ConnMaxIdleTime)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)

	return db, nil
}

func CloseDb(db *sqlx.DB, logger log.Logger) {
	err := db.Close()
	if err != nil {
		logger.Error(err, "failed closing database connection")
	}
}
