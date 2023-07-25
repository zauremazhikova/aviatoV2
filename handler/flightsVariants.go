package handler

import (
	"aviatoV2/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetAllFlights(c *fiber.Ctx) error {
	db := database.DB()
	if db != nil {
		rows, err := db.Query("SELECT ID, FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, COALESCE(SEATS_NUMBER, 0) AS SEATS_NUMBER, COALESCE(PRICE, 0) AS PRICE FROM flights")
		//var flightsMap fiber.Map

		//flights := []entities.Flight{}

		var (
			id            int
			flightNumber  string
			directionID   int
			departureTime time.Time
			arrivalTime   time.Time
			seatsNumber   int
			price         float64
		)

		for rows.Next() {
			//flight := entities.Flight{}
			err := rows.Scan(&id, &flightNumber, &directionID, &departureTime, &arrivalTime, &seatsNumber, &price)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(id, flightNumber, directionID, departureTime, arrivalTime, seatsNumber, price)
			//flights = append(flights, flight)
		}
		fmt.Println(err)
		/*if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not run query", "data": err})
		}
		defer rows.Close()
		err = rows.Err()
		if err != nil {
			return c.Status(204).JSON(fiber.Map{"status": "error", "message": "no content", "data": err})
		}
		return c.Status(200).JSON(fiber.Map{"status": "error", "message": "no content", "data": rows})*/

		db.Close()
	}
	//return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Could not connect to DB"})
	return nil
}
