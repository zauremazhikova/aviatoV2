package flight

import (
	"aviatoV2/entities/direction"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetAll(c *fiber.Ctx) error {

	responseFlight, err := GetAllFromDB()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flights not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": responseFlight})
}

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")

	responseFlight, err := GetSingleFromDB(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flight Found", "data": responseFlight})
}

func Create(c *fiber.Ctx) error {
	type insertStruct struct {
		FlightNumber  string    `json:"flightNumber"`
		DirectionID   string    `json:"directionID"`
		DepartureTime time.Time `json:"departureTime"`
		ArrivalTime   time.Time `json:"arrivalTime"`
		SeatsNumber   int       `json:"seatsNumber"`
		Price         float64   `json:"price"`
	}
	var insertData insertStruct
	err := c.BodyParser(&insertData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	currentDirection, _ := direction.GetSingleFromDB(insertData.DirectionID)
	if currentDirection.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Direction not found", "data": err})
	}

	flight := new(Flight)
	flight.FlightNumber = insertData.FlightNumber
	flight.Direction = *currentDirection
	flight.DepartureTime = insertData.DepartureTime
	flight.ArrivalTime = insertData.ArrivalTime
	flight.SeatsNumber = insertData.SeatsNumber
	flight.Price = insertData.Price

	err = CreateInDB(flight)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Flight", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flight has created"})
}

func Update(c *fiber.Ctx) error {
	type updateStruct struct {
		FlightNumber  string    `json:"flightNumber"`
		DirectionID   string    `json:"directionID"`
		DepartureTime time.Time `json:"departureTime"`
		ArrivalTime   time.Time `json:"arrivalTime"`
		SeatsNumber   int       `json:"seatsNumber"`
		Price         float64   `json:"price"`
	}
	id := c.Params("id")
	flight, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": err})
	}

	var updateData updateStruct
	err = c.BodyParser(&updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	currentDirection, _ := direction.GetSingleFromDB(updateData.DirectionID)
	if currentDirection.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Direction not found", "data": err})
	}

	flight.FlightNumber = updateData.FlightNumber
	flight.Direction = *currentDirection
	flight.DepartureTime = updateData.DepartureTime
	flight.ArrivalTime = updateData.ArrivalTime
	flight.SeatsNumber = updateData.SeatsNumber
	flight.Price = updateData.Price

	err = UpdateInDB(flight)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Flight has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flight has Updated", "data": flight})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	flight, err := GetSingleFromDB(id)

	if flight.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Flight", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flight deleted"})
}
