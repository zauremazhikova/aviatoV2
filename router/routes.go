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
	airlines := api.Group("/airline")
	country := api.Group("/country")
	/*	booking := api.Group("/booking")
		bookingStatus := api.Group("/bookingStatus")
		city := api.Group("/city")
		flight := api.Group("/flight")
		passenger := api.Group("/passenger")
		route := api.Group("/route")*/

	// main
	api.Get("/", handlers.GetAllFlights)

	// airline
	airlines.Get("/", airline.GetAll)
	airlines.Get("/:id", airline.GetSingle)
	airlines.Post("/", airline.Create)
	airlines.Put("/:id", airline.Update)
	airlines.Delete("/:id", airline.Delete)

	// country
	//country.Get("/", country.G)
	//country.Get("/:id", country.GetSingleCountry)
	//country.Post("/", country.CreateCountry)
	//country.Put("/:id", country.UpdateCountry)
	//country.Delete("/:id", country.DeleteCountryByID)

	/*


		// city
		city.Get("/", Entities.GetAllCities)
		city.Get("/:id", Entities.GetSingleCity)
		city.Post("/", Entities.CreateCity)
		city.Put("/:id", Entities.UpdateCity)
		city.Delete("/:id", Entities.DeleteCityByID)



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
