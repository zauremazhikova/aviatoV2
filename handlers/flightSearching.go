package handlers

import (
	"aviatoV2/config"
	"aviatoV2/entities/flight"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

var FlightsMap [][]*flight.Flight

func GetFlightsByOriginAndDestination(c *fiber.Ctx) error {
	type searchStruct struct {
		OriginCityID  string    `json:"originCityID"`
		DirectionID   string    `json:"destinationCityID"`
		DepartureTime time.Time `json:"departureTime"`
	}
	var searchData searchStruct
	err := c.BodyParser(&searchData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	flightsMap := make([][]*flight.Flight, 0)
	//flights := make([]*flight.Flight, 0)

	//FindFlightVariants(searchData.OriginCityID, searchData.DirectionID, config.FlightStopMaxNumber, flightsMap, flights)
	fmt.Println(flightsMap)

	return nil
}

func FindFlightVariants() {

	flights := make([]*flight.Flight, 0)
	res := findFlightVariants("1", "3", config.FlightStopMaxNumber, flights)
	fmt.Println(FlightsMap, res)

}

func findFlightVariants(originCityID string, destinationCityID string, stops int, flights []*flight.Flight) bool {

	if originCityID == destinationCityID {
		return true
	} else if stops <= 0 {
		return false
	}

	nextFlights, err := flight.GetFlightsByOriginCityFromDB(originCityID)
	if err != nil {
		return false
	}

	for i := 0; i < len(nextFlights); i++ {
		currentFlight := nextFlights[i]
		flights = append(flights, currentFlight)
		res := findFlightVariants(currentFlight.Direction.DestinationCity.ID, destinationCityID, stops-1, flights)
		if res == true {
			FlightsMap = append(FlightsMap, flights)
		}
	}
	return false

}
