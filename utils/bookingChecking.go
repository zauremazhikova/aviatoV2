package utils

import (
	"aviatoV2/entities/booking"
	"aviatoV2/entities/flight"
	"errors"
)

func CheckFlightBookingAvailability(flight *flight.Flight) (bool, error) {

	currentBookings, err := booking.GetByFlightIDFromDB(flight.ID)

	if err != nil {
		return false, err
	}

	if flight.SeatsNumber <= len(currentBookings) {
		return false, errors.New("flight is full")
	}

	return true, nil
}
