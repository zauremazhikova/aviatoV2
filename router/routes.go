package router

import (
	"aviatoV2/entities/airline"
	"aviatoV2/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	/*	booking := api.Group("/booking")
		bookingStatus := api.Group("/bookingStatus")
		city := api.Group("/city")
		country := api.Group("/country")
		flight := api.Group("/flight")
		passenger := api.Group("/passenger")
		route := api.Group("/route")*/

	// main
	api.Get("/", handlers.GetAllFlights)

	airlines := api.Group("/airline")
	airlines.Get("/:id", airline.GetSingleAirline)

	/*
		// airline
		airline.Get("/", Entities.GetAllAirlines)
		airline.Get("/:id", Entities.GetSingleAirline)
		airline.Post("/", Entities.CreateAirline)
		airline.Put("/:id", Entities.UpdateAirline)
		airline.Delete("/:id", Entities.DeleteAirlineByID)

		// city
		city.Get("/", Entities.GetAllCities)
		city.Get("/:id", Entities.GetSingleCity)
		city.Post("/", Entities.CreateCity)
		city.Put("/:id", Entities.UpdateCity)
		city.Delete("/:id", Entities.DeleteCityByID)

		// country
		country.Get("/", Entities.GetAllCountries)
		country.Get("/:id", Entities.GetSingleCountry)
		country.Post("/", Entities.CreateCountry)
		country.Put("/:id", Entities.UpdateCountry)
		country.Delete("/:id", Entities.DeleteCountryByID)

		// passenger
		passenger.Get("/", Entities.GetAllPassengers)
		passenger.Get("/:id", Entities.GetSinglePassenger)
		passenger.Post("/", Entities.CreatePassenger)
		passenger.Put("/:id", Entities.UpdatePassenger)
		passenger.Delete("/:id", Entities.DeletePassengerByID)

		// route
		route.Get("/", Entities.GetAllRoutes)
		route.Get("/:id", Entities.GetSingleRoute)
		route.Post("/", Entities.CreateRoute)
		route.Put("/:id", Entities.UpdateRoute)
		route.Delete("/:id", Entities.DeleteRouteByID)

		// flight
		flight.Get("/", Entities.GetAllFlights)
		flight.Get("/:id", Entities.GetSingleFlight)
		flight.Post("/", Entities.CreateFlight)
		flight.Put("/:id", Entities.UpdateFlight)
		flight.Delete("/:id", Entities.DeleteFlightByID)

		// booking
		booking.Get("/", Entities.GetAllBookings)
		booking.Get("/:id", Entities.GetSingleBooking)
		booking.Post("/", Entities.CreateBooking)
		booking.Put("/:id", Entities.UpdateBooking)
		booking.Delete("/:id", Entities.DeleteBookingByID)

		// booking Status
		bookingStatus.Get("/", Entities.GetAllBookingStatuses)
		bookingStatus.Get("/:id", Entities.GetSingleBookingStatus)
		bookingStatus.Post("/", Entities.CreateBookingStatus)
		bookingStatus.Put("/:id", Entities.UpdateBookingStatus)
		bookingStatus.Delete("/:id", Entities.DeleteBookingStatusByID)*/

}
