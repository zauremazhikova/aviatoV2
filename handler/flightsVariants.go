package handler

import (
	"aviatoV2/database"
	"aviatoV2/entities"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetAllFlights(c *fiber.Ctx) error {
	db := database.DB()
	if db != nil {
		rows, err := db.Query("SELECT * FROM flights")
		//var flightsMap fiber.Map

		flights := []entities.Flight{}

		for rows.Next() {
			flight := entities.Flight{}
			err := rows.Scan(&flight.ID, &flight.FlightNumber, &flight.Route, &flight.DepartureTime,
				&flight.ArrivalTime, &flight.SeatsNumber, &flight.Price)
			if err != nil {
				fmt.Println(err)
				continue
			}
			flights = append(flights, flight)
		}
		for _, flight := range flights {
			fmt.Println(flight.ID, flight.FlightNumber, flight.Route, flight.DepartureTime,
				flight.ArrivalTime, flight.SeatsNumber, flight.Price)
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
