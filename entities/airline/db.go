package airline

import (
	"aviatoV2/database"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func GetAllFromDB() (a []*Airline, err error) {
	airlines := make([]*Airline, 0)

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines")
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	for rows.Next() {
		var airline Airline
		err := rows.Scan(&airline.ID, &airline.Name, &airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt)
		if err != nil {
			return airlines, err
		} else {
			airlines = append(airlines, &airline)
		}
	}

	return airlines, nil
}

func GetSingleFromDB(id string) (*Airline, error) {

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines WHERE ID = $1", id)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	var airline Airline
	for rows.Next() {
		err := rows.Scan(&airline.ID, &airline.Name, &airline.CreatedAt, &airline.UpdatedAt, &airline.DeletedAt)
		if err != nil {
			return &Airline{}, err
		}
	}
	return &airline, nil

}

func CreateInDB(airline *Airline) error {
	db := database.DB()
	_, dbErr := db.Query("INSERT INTO airlines (name, created_at) VALUES ($1, $2)", airline.Name, time.Now())

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func UpdateInDB(airline *Airline) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE airlines SET name = $2, updated_at = $3 WHERE id = $1", airline.ID, airline.Name, time.Now())

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE airlines SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}
