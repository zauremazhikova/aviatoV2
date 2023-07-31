package handlers

import (
	"aviatoV2/config"
	"aviatoV2/entities/flight"
	"github.com/gofiber/fiber/v2"
	"slices"
	"time"
)

var flightsMap [][]*flight.Flight

func GetFlightsByOriginAndDestination(c *fiber.Ctx) error {

	type searchStruct struct {
		OriginCityID        string    `json:"OriginCityID"`
		DestinationCityID   string    `json:"DestinationCityID"`
		DepartureTime       time.Time `json:"departureTime"`
		FlightStopMaxNumber int       `json:"FlightStopMaxNumber"`
	}

	var searchData searchStruct
	err := c.BodyParser(&searchData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// maxStop - это максимальное количество пересадок. Настраивается в config. Но при поиске, пользователь может сам отрегулировать.
	// Если не регулирует: передается -1
	maxStop := config.FlightStopMaxNumber
	if searchData.FlightStopMaxNumber != -1 {
		maxStop = searchData.FlightStopMaxNumber
	}

	flightsMap = make([][]*flight.Flight, 0)
	findFlightsDFS(searchData.OriginCityID, searchData.DestinationCityID, maxStop, make([]*flight.Flight, 0), make([]string, 0))

	if len(flightsMap) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Flights not found", "data": flightsMap})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": flightsMap})
}

func findFlightsDFS(originCityID string, destinationCityID string, stops int, flights []*flight.Flight, citiesID []string) {

	contains := slices.Contains(citiesID, originCityID) // Проверка на то что город уже есть в списке. Чтобы избежать цикличных вариантов перелета. Например: Алматы -> Астана -> Алматы

	if originCityID == destinationCityID {
		flightsMap = append(flightsMap, flights)
		return
	} else if stops < 0 || contains == true {
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
		findFlightsDFS(currentFlight.Direction.DestinationCity.ID, destinationCityID, stops-1, flights, citiesID)
		flights = flights[:len(flights)-1]
	}
	return
}
