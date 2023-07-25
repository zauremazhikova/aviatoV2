package entities

import (
	"aviatoV2/database"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

type Flight struct {
	ID            int       `json:"id"`
	FlightNumber  string    `json:"flightNumber"`
	Direction     Direction `json:"direction"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	SeatsNumber   int       `json:"seatsNumber"`
	Price         float64   `json:"price"`
}

func CreateResponseFlight(id int, flightNumber string, directionID int, departureTime time.Time, arrivalTime time.Time, seatsNumber int, price float64) Flight {

	return Flight{
		ID:            id,
		FlightNumber:  flightNumber,
		Direction:     GetDirectionByID(directionID),
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		SeatsNumber:   seatsNumber,
		Price:         price}
}

func GetFlightByID(ID int) Flight {
	db := database.DB()
	rows, err := db.Query("SELECT ID, FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, SEATS_NUMBER, PRICE FROM flights WHERE ID = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id            int
		flightNumber  string
		directionID   int
		departureTime time.Time
		arrivalTime   time.Time
		seatsNumber   int
		price         float64
	)

	var city Flight
	for rows.Next() {
		err := rows.Scan(&id, &flightNumber, &directionID, &departureTime, &arrivalTime, &seatsNumber, &price)
		if err != nil {
			log.Fatal(err)
		}
		city = CreateResponseFlight(id, flightNumber, directionID, departureTime, arrivalTime, seatsNumber, price)

	}
	return city
}

/*
func GetFlightByID(id int) Flight {
	var flightModel models.Flight
	FindFlight(id, &flightModel)
	flight := CreateResponseFlight(flightModel)
	return flight
}

func GetFlightsByRoute(routeID int) []Flight {
	var flightsModel []models.Flight
	var flights []Flight
	FindFlightByRoute(routeID, &flightsModel)
	for _, flightModel := range flightsModel {
		flight := CreateResponseFlight(flightModel)
		flights = append(flights, flight)
	}
	return flights
}

// GetFlightByOriginCity from DB
func GetFlightByOriginCity(originCityID int) []Flight {
	routes := GetRoutesByOriginCity(originCityID)
	var flights []Flight

	for _, route := range routes {
		var flightsModel []models.Flight
		FindFlightByRoute(route.ID, &flightsModel)
		for _, flightModel := range flightsModel {
			flight := CreateResponseFlight(flightModel)
			flights = append(flights, flight)
		}
	}
	return flights
}

func GetFlights() []Flight {
	var flightsModel []models.Flight
	var flights []Flight
	FindAllFlight(&flightsModel)
	for _, flightModel := range flightsModel {
		flight := CreateResponseFlight(flightModel)
		flights = append(flights, flight)
	}
	return flights
}

// FindAllFlight from DB
func FindAllFlight(flights *[]models.Flight) {
	database.DB.Db.Find(&flights)
}

// FindFlight from DB
func FindFlight(id int, flight *models.Flight) {
	database.DB.Db.Find(&flight, "id = ?", id)
}

// FindFlightByRoute from DB
func FindFlightByRoute(routeID int, flight *[]models.Flight) {
	database.DB.Db.Find(&flight, "route_id = ?", routeID)
}

// Working with DB

// CreateFlight in DB
func CreateFlight(c *fiber.Ctx) error {
	db := database.DB.Db
	flight := new(models.Flight)

	err := c.BodyParser(flight)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&flight).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Flight", "data": err})
	}

	responseFlight := CreateResponseFlight(*flight)

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flight has created", "data": responseFlight})
}

// GetAllFlights from db
func GetAllFlights(c *fiber.Ctx) error {
	db := database.DB.Db
	var flights []models.Flight
	db.Find(&flights)

	if len(flights) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flights not found", "data": nil})
	}

	var responseFlights []Flight
	for _, flight := range flights {
		responseFlight := CreateResponseFlight(flight)
		responseFlights = append(responseFlights, responseFlight)
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Flights Found", "data": responseFlights})
}

// GetSingleFlight from db
func GetSingleFlight(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var flight models.Flight
	// find single flight in the database by id
	db.Find(&flight, "id = ?", id)
	if flight.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": nil})
	}

	responseFlight := CreateResponseFlight(flight)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Flight Found", "data": responseFlight})
}

// UpdateFlight in db
func UpdateFlight(c *fiber.Ctx) error {
	type updateFlight struct {
		FlightNumber  string    `json:"FlightNumber"`
		RouteID       int       `json:"RouteID"`
		DepartureTime time.Time `json:"departureTime"`
		ArrivalTime   time.Time `json:"arrivalTime"`
		SeatsNumber   int       `json:"seatsNumber"`
		Price         float64   `json:"price"`
	}
	db := database.DB.Db
	var flight models.Flight
	// get id params
	id := c.Params("id")
	db.Find(&flight, "id = ?", id)
	if flight.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": nil})
	}
	var updateFlightData updateFlight
	err := c.BodyParser(&updateFlightData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	flight.FlightNumber = updateFlightData.FlightNumber
	flight.RouteID = updateFlightData.RouteID
	flight.DepartureTime = updateFlightData.DepartureTime
	flight.ArrivalTime = updateFlightData.ArrivalTime
	flight.SeatsNumber = updateFlightData.SeatsNumber
	flight.Price = updateFlightData.Price

	db.Save(&flight)

	responseFlight := CreateResponseFlight(flight)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "flight Found", "data": responseFlight})
}

// DeleteFlightByID in db
func DeleteFlightByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var flight models.Flight
	// get id params
	id := c.Params("id")
	// find single flight in the database by id
	db.Find(&flight, "id = ?", id)
	if flight.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": nil})
	}
	err := db.Delete(&flight, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete flight", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Flight deleted"})
}
*/
