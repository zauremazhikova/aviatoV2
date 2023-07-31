package booking

import (
	"aviatoV2/database"
	"aviatoV2/entities/flight"
	"aviatoV2/entities/passenger"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func GetAllFromDB() (a []*Booking, err error) {
	bookings := make([]*Booking, 0)

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, BOOKING_NUMBER, FLIGHT_ID, PASSENGER_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM bookings")
	if dbErr != nil {
		db.Close()
		log.Fatal(dbErr)
	}

	for rows.Next() {
		var booking Booking
		var flightID string
		var passengerID string
		err := rows.Scan(&booking.ID, &booking.BookingNumber, &flightID, &passengerID, &booking.CreatedAt, &booking.UpdatedAt, &booking.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return bookings, err
		} else {
			currentFlight, _ := flight.GetSingleFromDB(flightID)
			booking.Flight = *currentFlight

			currentPassenger, _ := passenger.GetSingleFromDB(passengerID)
			booking.Passenger = *currentPassenger

			bookings = append(bookings, &booking)
		}
	}
	db.Close()
	return bookings, nil
}

func GetByFlightIDFromDB(flightID string) (a []*Booking, err error) {
	bookings := make([]*Booking, 0)

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, BOOKING_NUMBER, FLIGHT_ID, PASSENGER_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM bookings WHERE FLIGHT_ID = $1", flightID)
	if dbErr != nil {
		db.Close()
		log.Fatal(dbErr)
	}

	for rows.Next() {
		var booking Booking
		var flightID string
		var passengerID string
		err := rows.Scan(&booking.ID, &booking.BookingNumber, &flightID, &passengerID, &booking.CreatedAt, &booking.UpdatedAt, &booking.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return bookings, err
		} else {
			currentFlight, _ := flight.GetSingleFromDB(flightID)
			booking.Flight = *currentFlight

			currentPassenger, _ := passenger.GetSingleFromDB(passengerID)
			booking.Passenger = *currentPassenger

			bookings = append(bookings, &booking)
		}
	}
	db.Close()
	return bookings, nil
}

func GetSingleFromDB(id string) (*Booking, error) {

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, BOOKING_NUMBER, FLIGHT_ID, PASSENGER_ID, CREATED_AT, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM bookings WHERE ID = $1", id)
	if dbErr != nil {
		db.Close()
		log.Fatal(dbErr)
	}

	var booking Booking
	var flightID string
	var passengerID string

	for rows.Next() {
		err := rows.Scan(&booking.ID, &booking.BookingNumber, &flightID, &passengerID, &booking.CreatedAt, &booking.UpdatedAt, &booking.DeletedAt)
		if err != nil {
			return &Booking{}, err
		}
	}
	currentFlight, _ := flight.GetSingleFromDB(flightID)
	booking.Flight = *currentFlight

	currentPassenger, _ := passenger.GetSingleFromDB(passengerID)
	booking.Passenger = *currentPassenger

	db.Close()
	return &booking, nil

}

func CreateInDB(booking *Booking) error {
	db := database.DB()
	_, dbErr := db.Query("INSERT INTO bookings (booking_number, flight_id, passenger_id, created_at) VALUES ($1, $2, $3, $4)", booking.BookingNumber, booking.Flight.ID, booking.Passenger.ID, time.Now())

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}

func UpdateInDB(booking *Booking) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE bookings SET booking_number = $2, flight_id = $3, passenger_id = $4, updated_at = $3 WHERE id = $1",
		booking.ID, booking.BookingNumber, booking.Flight.ID, booking.Passenger.ID, time.Now())

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE bookings SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}
