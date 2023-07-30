package flight

import (
	"aviatoV2/database"
	"aviatoV2/entities/direction"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func GetAllFromDB() (a []*Flight, err error) {
	flights := make([]*Flight, 0)

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, SEATS_NUMBER, PRICE, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM flights")
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	for rows.Next() {
		var flight Flight
		var directionID string
		err := rows.Scan(&flight.ID, &flight.FlightNumber, &directionID, &flight.DepartureTime, &flight.ArrivalTime, &flight.SeatsNumber, &flight.Price, &flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt)
		if err != nil {
			return flights, err
		} else {
			currDirection, _ := direction.GetSingleFromDB(directionID)
			flight.Direction = *currDirection
			flights = append(flights, &flight)
		}
	}

	return flights, nil
}

func GetSingleFromDB(id string) (*Flight, error) {

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, SEATS_NUMBER, PRICE, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM flights WHERE ID = $1", id)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	var flight Flight
	var directionID string
	for rows.Next() {
		err := rows.Scan(&flight.ID, &flight.FlightNumber, &directionID, &flight.DepartureTime, &flight.ArrivalTime, &flight.SeatsNumber, &flight.Price, &flight.CreatedAt, &flight.UpdatedAt, &flight.DeletedAt)
		if err != nil {
			return &Flight{}, err
		}
	}
	currDirection, _ := direction.GetSingleFromDB(directionID)
	flight.Direction = *currDirection
	return &flight, nil

}

func CreateInDB(flight *Flight) error {
	db := database.DB()
	_, dbErr := db.Query("INSERT INTO flights (FLIGHT_NUMBER, DIRECTION_ID, DEPARTURE_TIME, ARRIVAL_TIME, SEATS_NUMBER, PRICE, CREATED_AT) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		flight.FlightNumber, flight.Direction.ID, flight.DepartureTime, flight.ArrivalTime, flight.SeatsNumber, flight.Price, time.Now())

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func UpdateInDB(flight *Flight) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE flights SET FLIGHT_NUMBER = $2, DIRECTION_ID = $3, DEPARTURE_TIME = $4, ARRIVAL_TIME = $5, SEATS_NUMBER = $6, PRICE = $7, UPDATED_AT = $8 WHERE id = $1",
		flight.ID, flight.FlightNumber, flight.Direction.ID, flight.DepartureTime, flight.ArrivalTime, flight.SeatsNumber, flight.Price, time.Now())

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE flights SET DELETED_AT = $1 WHERE id = $2", time.Now(), id)

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}
