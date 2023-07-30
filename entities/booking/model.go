package booking

import (
	"aviatoV2/entities/flight"
	"aviatoV2/entities/passenger"
	"time"
)

type Booking struct {
	ID            string              `json:"id"`
	BookingNumber string              `json:"bookingNumber"`
	Flight        flight.Flight       `json:"flight"`
	Passenger     passenger.Passenger `json:"passenger"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	DeletedAt     time.Time           `json:"deleted_at"`
}
