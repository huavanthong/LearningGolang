package main

import (
	"github.com/gofiber/fiber"

	"github.com/huavanthong/MasterGolang/tree/main/03_ReferenceProject/09_docker-compose-go-app/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

// Setup - set's up our fiber app and the routes
// returns a pointer to app
func SetUp() *fiber.App {
	app := fiber.New()
	setupRoutes(app)
	return app
}
