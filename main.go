package main

import "aviatoV2/handlers"

func main() {
	/*fiber.New()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}*/

	handlers.FindFlightVariants()
}
