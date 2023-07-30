package handlers

import (
	"aviatoV2/config"
	"aviatoV2/entities/flight"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"slices"
)

var FlightsMap [][]*flight.Flight

func GetFlightsByOriginAndDestination(c *fiber.Ctx) error {
	FlightsMap = make([][]*flight.Flight, 0)
	type searchStruct struct {
		OriginCityID string `json:"originCityID"`
		DirectionID  string `json:"destinationCityID"`
	}
	var searchData searchStruct
	err := c.BodyParser(&searchData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	findFlightVariants(searchData.OriginCityID, searchData.DirectionID, config.FlightStopMaxNumber, make([]*flight.Flight, 0), make([]string, 0))

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": FlightsMap})
}

func FindFlightVariants() {
	findFlightVariants("1", "3", config.FlightStopMaxNumber, make([]*flight.Flight, 0), make([]string, 0))
	fmt.Println(FlightsMap)
}

func findFlightVariants(originCityID string, destinationCityID string, stops int, flights []*flight.Flight, citiesID []string) {

	contains := slices.Contains(citiesID, originCityID)

	if originCityID == destinationCityID {
		FlightsMap = append(FlightsMap, flights)
		return
	} else if stops <= 0 || contains == true {
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
