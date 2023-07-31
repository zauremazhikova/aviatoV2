package booking

import (
	"aviatoV2/entities/flight"
	"aviatoV2/entities/passenger"
	"aviatoV2/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	responseBooking, err := GetAllFromDB()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Bookings not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Bookings Found", "data": responseBooking})
}

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")

	responseBooking, err := GetSingleFromDB(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Booking not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Booking Found", "data": responseBooking})
}

func Create(c *fiber.Ctx) error {
	type insertStruct struct {
		BookingNumber string `json:"bookingNumber"`
		FlightID      string `json:"flightID"`
		PassengerID   string `json:"passengerID"`
	}
	var insertData insertStruct
	err := c.BodyParser(&insertData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	currentFlight, _ := flight.GetSingleFromDB(insertData.FlightID)
	if currentFlight.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": err})
	}

	checkingAvailability, err := utils.CheckFlightBookingAvailability(currentFlight)
	if checkingAvailability == false {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight is not available", "data": err})
	}

	currentPassenger, err := passenger.GetSingleFromDB(insertData.PassengerID)
	if currentPassenger.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": err})
	}

	booking := new(Booking)
	booking.BookingNumber = insertData.BookingNumber
	booking.Flight = *currentFlight
	booking.Passenger = *currentPassenger

	err = CreateInDB(booking)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Booking", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Booking has created"})
}

func Update(c *fiber.Ctx) error {
	type updateStruct struct {
		BookingNumber string `json:"bookingNumber"`
		FlightID      string `json:"flightID"`
		PassengerID   string `json:"passengerID"`
	}
	id := c.Params("id")
	booking, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Booking not found", "data": err})
	}

	var updateData updateStruct
	err = c.BodyParser(&updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	currentFlight, err := flight.GetSingleFromDB(updateData.FlightID)
	if currentFlight.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flight not found", "data": err})
	}

	currentPassenger, err := passenger.GetSingleFromDB(updateData.PassengerID)
	if currentPassenger.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": err})
	}

	booking.BookingNumber = updateData.BookingNumber
	booking.Flight = *currentFlight
	booking.Passenger = *currentPassenger

	err = UpdateInDB(booking)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Booking has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Booking has Updated", "data": booking})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	booking, err := GetSingleFromDB(id)

	if booking.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Booking not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Booking", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Booking deleted"})
}
