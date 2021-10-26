package main

import (
	"fmt"
	"net/http"

	tHttp "github.com/mohammadraasel/go-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	handler := tHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":4000", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our API")
	}
}
