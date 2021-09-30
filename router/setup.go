package router

import (
	"github.com/gofiber/fiber/v2"
	db "movies/database"
	"movies/models"
)

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("hello", func(c *fiber.Ctx) error { return c.SendString("hello world") })
	api.Get("movies", getMovies)
	api.Get("directors", getDirectors)
	api.Post("movies", addMovie)
	api.Post("directors", addDirector)
}

func getMovies(c *fiber.Ctx) error {
	var movies []models.Movie
	db.DB.Find(&movies)
	return c.JSON(movies)
}

func getDirectors(c *fiber.Ctx) error {
	var directors []models.Director
	db.DB.Find(&directors)
	return c.JSON(directors)
}

func addMovie(c *fiber.Ctx) error {
	movie := new(models.Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	if err := db.DB.Create(&movie).Error; err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later. 😕",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(movie)
}

func addDirector(c *fiber.Ctx) error {

	director := new(models.Director)
	if err := c.BodyParser(director); err != nil {
		return c.JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	if err := db.DB.Create(&director).Error; err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later. 😕",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(director)
}
