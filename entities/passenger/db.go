package passenger

import (
	"aviatoV2/database"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func GetAllFromDB() (a []*Passenger, err error) {
	countries := make([]*Passenger, 0)

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, PASSPORT, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM passengers")
	if dbErr != nil {
		db.Close()
		log.Fatal(dbErr)
	}

	for rows.Next() {
		var passenger Passenger
		err := rows.Scan(&passenger.ID, &passenger.Name, &passenger.Passport, &passenger.CreatedAt, &passenger.UpdatedAt, &passenger.DeletedAt)
		if err != nil {
			return countries, err
		} else {
			countries = append(countries, &passenger)
		}
	}

	db.Close()
	return countries, nil
}

func GetSingleFromDB(id string) (*Passenger, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, PASSPORT, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM passengers WHERE ID = $1", id)
	if err != nil {
		db.Close()
		return &Passenger{}, err
	}

	var passenger Passenger
	for rows.Next() {
		err := rows.Scan(&passenger.ID, &passenger.Name, &passenger.Passport, &passenger.CreatedAt, &passenger.UpdatedAt, &passenger.DeletedAt)
		if err != nil {
			return &Passenger{}, err
		}
	}

	db.Close()
	return &passenger, nil
}

func CreateInDB(passenger *Passenger) error {
	db := database.DB()
	_, dbErr := db.Query("INSERT INTO passengers (name, passport, created_at) VALUES ($1, $2, $3)", passenger.Name, passenger.Passport, time.Now())

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}

func UpdateInDB(passenger *Passenger) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE passengers SET name = $2, passport = $3, updated_at = $4 WHERE id = $1", passenger.ID, passenger.Name, passenger.Passport, time.Now())

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE passengers SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}
