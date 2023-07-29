package main

import (
	"aviatoV2/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fiber.New()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
	//a, err := airline.FindOne("4")
	//fmt.Println(a, err)
}
