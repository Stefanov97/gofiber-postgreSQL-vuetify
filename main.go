package main

import (
	"log"
	"movies/database"
	"movies/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to Postgres
	database.ConnectToDB()
	app := fiber.New()
	router.SetupRoutes(app)
	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
