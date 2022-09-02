//go:build wireinject

package main

import (
	"awesomeProject1/api"
	"awesomeProject1/config"
	"awesomeProject1/db"
	"awesomeProject1/handlers"
	"awesomeProject1/log"
	"awesomeProject1/repository"
	"awesomeProject1/service"
	"github.com/google/wire"
)

func InitializeApp() (App, error) {
	wire.Build(
		// app
		NewApp,

		// api
		api.NewApi,

		// handlers
		handlers.NewGetConfigHandler,
		handlers.NewGetUsersHandler,
		handlers.NewPostUsersHandler,
		handlers.NewEditUserHandler,

		// config
		config.NewConfig,

		// logger
		log.NewZeroLogger,
		wire.Bind(new(log.Logger), new(log.ZeroLogger)),

		// services
		service.NewPostgresUserService,
		wire.Bind(new(service.UserService), new(service.PostgresUserService)),

		// repositories
		repository.NewPostgresUserRepository,
		wire.Bind(new(repository.UserRepository), new(repository.PostgresUserRepository)),

		// database
		db.NewDb,
	)

	return App{}, nil
}
