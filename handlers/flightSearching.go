package handlers

import (
	"aviatoV2/config"
	"aviatoV2/entities/flight"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"slices"
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

	findFlightVariants("1", "3", config.FlightStopMaxNumber, make([]*flight.Flight, 0), make([]string, 0))

	return nil
}

func FindFlightVariants() {
	findFlightVariants("1", "3", config.FlightStopMaxNumber, make([]*flight.Flight, 0), make([]string, 0))
	fmt.Println(FlightsMap)
}

func findFlightVariants(originCityID string, destinationCityID string, stops int, flights []*flight.Flight, citiesID []string) {

	contains := slices.Contains(citiesID, originCityID)
	correctTime := true
	n := len(flights)
	if n > 1 {
		flight1 := flights[n-1]
		flight2 := flights[n-2]
		if flight1.DepartureTime.Before(flight2.ArrivalTime) {
			correctTime = false
		} else {
			diff := flight1.DepartureTime.Sub(flight2.ArrivalTime)
			fmt.Println(diff)
		}
	}

	if originCityID == destinationCityID {
		FlightsMap = append(FlightsMap, flights)
		return
	} else if stops <= 0 || contains == true || correctTime == false {
		flights = nil
		return
	}

	nextFlights, err := flight.GetFlightsByOriginCityFromDB(originCityID)
	if err != nil {
		return
	}

	for i := 0; i < len(nextFlights); i++ {
		currentFlight := nextFlights[i]
		flights = append(flights, currentFlight)
		citiesID = append(citiesID, originCityID)
		findFlightVariants(currentFlight.Direction.DestinationCity.ID, destinationCityID, stops-1, flights, citiesID)
		flights = flights[:len(flights)-1]
	}
	return
}
