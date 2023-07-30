package handlers

/*
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
	findFlightVariants("1", "3", config.FlightStopMaxNumber, flights)
	fmt.Println(FlightsMap)
}

func findFlightVariants(originCityID string, destinationCityID string, stops int, flights []*flight.Flight) {

		currOriginCity, _ := flight.GetSingleFromDB(originCityID)
		flights = append(flights, currOriginCity)

		if originCityID == destinationCityID {
			FlightsMap = append(FlightsMap, flights)
			return
		} else if stops <= 0 {
			return
		}

		nextFlights, err := flight.GetFlightsByOriginCityFromDB(originCityID)
		if err != nil {
			return
		}

		for i := 0; i < len(nextFlights); i++ {
			findFlightVariants(nextFlights[i].Direction.DestinationCity.ID, destinationCityID, stops-1, flights)
		}*
	return
}*/
