package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Redis    Redis
	Database Database
}

type Redis struct {
	Host     string
	Password string
}

type Database struct {
	Driver          string
	Datasource      string
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

const redisHostKey = "REDIS_HOST"
const redisPortKey = "REDIS_PORT"
const databaseDriverKey = "DATABASE_DRIVER"
const databaseDatasourceKey = "DATABASE_DATASOURCE"
const databaseConnMaxLifetimeKey = "DATABASE_CONN_MAX_LIFETIME"
const databaseConnMaxIdleTimeKey = "DATABASE_CONN_MAX_IDLE_TIME"
const databaseMaxOpenConnsKey = "DATABASE_MAX_OPEN_CONNS"
const databaseMaxIdleConnsKey = "DATABASE_MAX_IDLE_CONNS"

func NewConfig() Config {
	// search for env variables that start with "AP" prefix
	viper.SetEnvPrefix("AP")

	// set default values if no environment variables present
	viper.SetDefault(databaseConnMaxLifetimeKey, 0)
	viper.SetDefault(databaseConnMaxIdleTimeKey, 0)
	viper.SetDefault(databaseMaxOpenConnsKey, 0)
	viper.SetDefault(databaseMaxIdleConnsKey, 0)

	// load environment variables into Viper
	viper.AutomaticEnv()

	// populate our Config struct
	cfg := Config{
		Redis: Redis{
			Host:     viper.GetString(redisHostKey),
			Password: viper.GetString(redisPortKey),
		},
		Database: Database{
			Driver:          viper.GetString(databaseDriverKey),
			Datasource:      viper.GetString(databaseDatasourceKey),
			ConnMaxLifetime: viper.GetDuration(databaseConnMaxLifetimeKey),
			ConnMaxIdleTime: viper.GetDuration(databaseConnMaxIdleTimeKey),
			MaxOpenConns:    viper.GetInt(databaseMaxOpenConnsKey),
			MaxIdleConns:    viper.GetInt(databaseMaxIdleConnsKey),
		},
	}

	return cfg
}
