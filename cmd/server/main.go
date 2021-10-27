package main

import (
	"fmt"
	"net/http"

	"github.com/mohammadraasel/go-api/internal/comment"
	"github.com/mohammadraasel/go-api/internal/database"
	tHttp "github.com/mohammadraasel/go-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {

	db, err := database.New()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)
	handler := tHttp.NewHandler(commentService)
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
