package entities

import (
	"aviatoV2/database"
	"github.com/gofiber/fiber/v2/log"
)

type Booking struct {
	ID            int       `json:"id"`
	BookingNumber string    `json:"bookingNumber"`
	Flight        Flight    `json:"flight"`
	Passenger     Passenger `json:"passenger"`
}

func CreateResponseBooking(id int, bookingNumber string, flightID int, passengerID int) Booking {

	return Booking{
		ID:            id,
		BookingNumber: bookingNumber,
		Flight:        GetFlightByID(flightID),
		Passenger:     GetPassengerByID(passengerID),
	}
}

func GetBookingByID(ID int) Booking {

	db := database.DB()
	rows, err := db.Query("SELECT ID, BOOKING_NUMBER, FLIGHT_ID, PASSENGER_ID FROM bookings WHERE ID = ?", ID)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id            int
		bookingNumber string
		flightID      int
		passengerID   int
	)

	var booking Booking

	for rows.Next() {
		err := rows.Scan(&id, &bookingNumber, &flightID, &passengerID)
		if err != nil {
			log.Fatal(err)
		}
		booking = CreateResponseBooking(id, bookingNumber, flightID, passengerID)

	}
	return booking
}

/*


// CreateBooking in DB
func CreateBooking(c *fiber.Ctx) error {
	db := database.DB.Db
	booking := new(models.Booking)

	err := c.BodyParser(booking)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	err = db.Create(&booking).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Booking", "data": err})
	}

	responseBooking := CreateResponseBooking(*booking)

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Booking has created", "data": responseBooking})
}

// GetAllBookings from db
func GetAllBookings(c *fiber.Ctx) error {
	db := database.DB.Db
	var bookings []models.Booking

	db.Find(&bookings)

	if len(bookings) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Bookings not found", "data": nil})
	}

	var responseBookings []Booking
	for _, booking := range bookings {
		responseBooking := CreateResponseBooking(booking)
		responseBookings = append(responseBookings, responseBooking)
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Bookings Found", "data": responseBookings})
}

// FindBooking by ID
func FindBooking(id int, booking *models.Booking) {
	database.DB.Db.Find(&booking, "id = ?", id)
}

// GetSingleBooking from db
func GetSingleBooking(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")
	var booking models.Booking

	db.Find(&booking, "id = ?", id)
	if booking.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Booking not found", "data": nil})
	}

	responseBooking := CreateResponseBooking(booking)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Booking Found", "data": responseBooking})
}

// UpdateBooking in db
func UpdateBooking(c *fiber.Ctx) error {
	type updateBooking struct {
		BookingNumber string `json:"BookingNumber"`
		FlightID      int    `json:"FlightID"`
		PassengerID   int    `json:"PassengerID"`
	}
	db := database.DB.Db
	var booking models.Booking

	id := c.Params("id")

	db.Find(&booking, "id = ?", id)
	if booking.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Booking not found", "data": nil})
	}
	var updateBookingData updateBooking
	err := c.BodyParser(&updateBookingData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	booking.BookingNumber = updateBookingData.BookingNumber
	booking.FlightID = updateBookingData.FlightID
	booking.PassengerID = updateBookingData.PassengerID
	db.Save(&booking)

	responseBooking := CreateResponseBooking(booking)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "booking Found", "data": responseBooking})
}

// DeleteBookingByID in db
func DeleteBookingByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var booking models.Booking

	id := c.Params("id")

	db.Find(&booking, "id = ?", id)
	if booking.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Booking not found", "data": nil})
	}
	err := db.Delete(&booking, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete Booking", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Booking deleted"})
}
*/
