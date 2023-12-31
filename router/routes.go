package router

import (
	"aviatoV2/entities/airline"
	"aviatoV2/entities/booking"
	"aviatoV2/entities/city"
	"aviatoV2/entities/country"
	"aviatoV2/entities/direction"
	"aviatoV2/entities/flight"
	"aviatoV2/entities/passenger"
	"aviatoV2/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	airlines := api.Group("/airline")
	bookings := api.Group("/booking")
	cities := api.Group("/city")
	countries := api.Group("/country")
	directions := api.Group("/direction")
	flights := api.Group("/flight")
	passengers := api.Group("/passenger")

	// main
	api.Get("/", handlers.GetFlightsByOriginAndDestination)

	// airline
	airlines.Get("/", airline.GetAll)
	airlines.Get("/:id", airline.GetSingle)
	airlines.Post("/", airline.Create)
	airlines.Put("/:id", airline.Update)
	airlines.Delete("/:id", airline.Delete)

	// booking
	bookings.Get("/", booking.GetAll)
	bookings.Get("/:id", booking.GetSingle)
	bookings.Post("/", booking.Create)
	bookings.Put("/:id", booking.Update)
	bookings.Delete("/:id", booking.Delete)

	// city
	cities.Get("/", city.GetAll)
	cities.Get("/:id", city.GetSingle)
	cities.Post("/", city.Create)
	cities.Put("/:id", city.Update)
	cities.Delete("/:id", city.Delete)

	// country
	countries.Get("/", country.GetAll)
	countries.Get("/:id", country.GetSingle)
	countries.Post("/", country.Create)
	countries.Put("/:id", country.Update)
	countries.Delete("/:id", country.Delete)

	// direction
	directions.Get("/", direction.GetAll)
	directions.Get("/:id", direction.GetSingle)
	directions.Post("/", direction.Create)
	directions.Put("/:id", direction.Update)
	directions.Delete("/:id", direction.Delete)

	// flight
	flights.Get("/", flight.GetAll)
	flights.Get("/:id", flight.GetSingle)
	flights.Post("/", flight.Create)
	flights.Put("/:id", flight.Update)
	flights.Delete("/:id", flight.Delete)

	// passenger
	passengers.Get("/", passenger.GetAll)
	passengers.Get("/:id", passenger.GetSingle)
	passengers.Post("/", passenger.Create)
	passengers.Put("/:id", passenger.Update)
	passengers.Delete("/:id", passenger.Delete)

}
