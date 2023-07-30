package direction

import (
	"aviatoV2/entities/airline"
	"aviatoV2/entities/city"
	"time"
)

type Direction struct {
	ID              int             `json:"id"`
	OriginCity      city.City       `json:"originCity"`
	DestinationCity city.City       `json:"destinationCity"`
	Airline         airline.Airline `json:"airline"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `json:"deleted_at"`
}
