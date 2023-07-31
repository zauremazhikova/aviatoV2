package handlers

import (
	"aviatoV2/config"
	"aviatoV2/entities/flight"
	"github.com/gofiber/fiber/v2"
	"slices"
)

var flightsMap [][]*flight.Flight

func GetFlightsByOriginAndDestination(c *fiber.Ctx) error {

	type searchStruct struct {
		OriginCityID      string `json:"OriginCityID"`
		DestinationCityID string `json:"DestinationCityID"`
	}

	var searchData searchStruct
	err := c.BodyParser(&searchData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	flightsMap = make([][]*flight.Flight, 0)
	findFlightVariants(searchData.OriginCityID, searchData.DestinationCityID, config.FlightStopMaxNumber, make([]*flight.Flight, 0), make([]string, 0))

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": flightsMap})
}

func findFlightVariants(originCityID string, destinationCityID string, stops int, flights []*flight.Flight, citiesID []string) {

	contains := slices.Contains(citiesID, originCityID)

	if originCityID == destinationCityID {
		flightsMap = append(flightsMap, flights)
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
