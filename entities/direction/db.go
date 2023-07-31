package direction

import (
	"aviatoV2/database"
	"aviatoV2/entities/airline"
	"aviatoV2/entities/city"
	"fmt"
	"time"
)

func GetAllFromDB() (a []*Direction, err error) {
	directions := make([]*Direction, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, ORIGIN_CITY_ID, DESTINATION_CITY_ID, AIRLINE_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM directions")
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var direction Direction
		var originCityID string
		var destinationCityID string
		var airlineCityID string

		err := rows.Scan(&direction.ID, &originCityID, &destinationCityID, &airlineCityID, &direction.CreatedAt, &direction.UpdatedAt, &direction.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return directions, err
		} else {
			currOriginCity, _ := city.GetSingleFromDB(originCityID)
			direction.OriginCity = *currOriginCity

			currDestCity, _ := city.GetSingleFromDB(destinationCityID)
			direction.DestinationCity = *currDestCity

			currAirline, _ := airline.GetSingleFromDB(airlineCityID)
			direction.Airline = *currAirline

			directions = append(directions, &direction)
		}
	}

	return directions, nil
}

func GetSingleFromDB(id string) (*Direction, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, ORIGIN_CITY_ID, DESTINATION_CITY_ID, AIRLINE_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM directions WHERE ID = $1", id)
	_ = db.Close()
	if err != nil {
		return nil, err
	}

	var direction Direction
	var originCityID string
	var destinationCityID string
	var airlineCityID string

	for rows.Next() {
		err := rows.Scan(&direction.ID, &originCityID, &destinationCityID, &airlineCityID, &direction.CreatedAt, &direction.UpdatedAt, &direction.DeletedAt)
		if err != nil {
			return &Direction{}, err
		}
	}
	currOriginCity, _ := city.GetSingleFromDB(originCityID)
	direction.OriginCity = *currOriginCity

	currDestCity, _ := city.GetSingleFromDB(destinationCityID)
	direction.DestinationCity = *currDestCity

	currAirline, _ := airline.GetSingleFromDB(airlineCityID)
	direction.Airline = *currAirline

	return &direction, nil
}

func CreateInDB(direction *Direction) error {
	db := database.DB()
	_, err := db.Query("INSERT INTO directions (origin_city_id, destination_city_id, airline_id, created_at) VALUES ($1, $2, $3, $4)", direction.OriginCity.ID, direction.DestinationCity.ID, direction.Airline.ID, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdateInDB(direction *Direction) error {
	db := database.DB()
	_, err := db.Query("UPDATE directions SET origin_city_id = $2, destination_city_id = $3, airline_id = $4, updated_at = $5 WHERE id = $1", direction.ID, direction.OriginCity.ID, direction.DestinationCity.ID, direction.Airline.ID, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, err := db.Query("UPDATE directions SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}
