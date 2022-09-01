package main

import (
	"awesomeProject1/api"
)

type App struct {
	Api api.Api
}

func (app App) Run() error {
	return app.Api.Startup()
}

func NewApp(api api.Api) App {
	return App{
		Api: api,
	}
}
