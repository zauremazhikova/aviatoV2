package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllFlights(c *fiber.Ctx) error {
	/*
		db := database.DB()
		if db != nil {
			rows, err := db.Query("" +
				"SELECT ID, FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, COALESCE(SEATS_NUMBER, 0) AS SEATS_NUMBER, COALESCE(PRICE, 0) AS PRICE, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM flights")

			var (
				id            int
				flightNumber  string
				directionID   int
				departureTime time.Time
				arrivalTime   time.Time
				seatsNumber   int
				price         float64
				createdAt     time.Time
				updatedAt     time.Time
				deletedAt     time.Time
			)

			var responseFlights []entities.Flight

			for rows.Next() {
				err := rows.Scan(&id, &flightNumber, &directionID, &departureTime, &arrivalTime, &seatsNumber, &price, &createdAt, &updatedAt, &deletedAt)
				if err != nil {
					fmt.Println(err)
					continue
				}
				//responseFlight := entities.CreateResponseFlight(id, flightNumber, directionID, departureTime, arrivalTime, seatsNumber, price, createdAt, updatedAt, deletedAt)
				responseFlight := entities.GetFlightByID(id)
				responseFlights = append(responseFlights, responseFlight)
				//fmt.Println(id, flightNumber, directionID, departureTime, arrivalTime, seatsNumber, price)
			}
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not run query", "data": err})
			}
			defer rows.Close()
			err = rows.Err()
			if err != nil {
				return c.Status(204).JSON(fiber.Map{"status": "error", "message": "no content", "data": err})
			}
			return c.Status(200).JSON(fiber.Map{"status": "error", "message": "no content", "data": responseFlights})

			db.Close()
		}
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Could not connect to DB"})*/
	return nil
}
