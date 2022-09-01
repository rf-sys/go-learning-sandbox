package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
