package country

import (
	"aviatoV2/database"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func GetAllFromDB() (a []*Country, err error) {
	countries := make([]*Country, 0)

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM countries")
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	for rows.Next() {
		var country Country
		err := rows.Scan(&country.ID, &country.Name, &country.CreatedAt, &country.UpdatedAt, &country.DeletedAt)
		if err != nil {
			return countries, err
		} else {
			countries = append(countries, &country)
		}
	}

	return countries, nil
}

func GetSingleFromDB(id string) (*Country, error) {

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM countries WHERE ID = $1", id)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	var country Country
	for rows.Next() {
		err := rows.Scan(&country.ID, &country.Name, &country.CreatedAt, &country.UpdatedAt, &country.DeletedAt)
		if err != nil {
			return &Country{}, err
		}
	}
	return &country, nil

}

func CreateInDB(country *Country) error {
	db := database.DB()
	_, dbErr := db.Query("INSERT INTO countries (name, created_at) VALUES ($1, $2)", country.Name, time.Now())

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func UpdateInDB(country *Country) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE countries SET name = $1, updated_at = $2 WHERE id = $3", country.Name, time.Now(), country.ID)

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE countries SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}
