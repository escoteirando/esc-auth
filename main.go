package main

import (
	"github.com/escoteirando/esc-auth/internal/controllers"
	"github.com/escoteirando/esc-auth/migrations"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// Add migrations to run
	app.Migrate(migrations.All(app.Config))

	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})
	app.POST("/login", controllers.LoginHandler)

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
