package airline

import (
	"aviatoV2/database"
	"github.com/gofiber/fiber/v2/log"
)

func Create(airline *Airline) error {
	return nil
}

func FindAll() (u []Airline, err error) {
	authors := make([]Airline, 0)
	return authors, nil
}

func FindOne(airlineID string) (*Airline, error) {

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines WHERE ID = $1", airlineID)
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

func Update(airline Airline) error {
	return nil
}

func Delete(id int) error {
	return nil
}
