package city

import (
	"aviatoV2/entities/country"
	"time"
)

type City struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Country   country.Country `json:"country"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt time.Time       `json:"deleted_at"`
}
