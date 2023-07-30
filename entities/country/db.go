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
		db.Close()
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

	db.Close()
	return countries, nil
}

func GetSingleFromDB(id string) (*Country, error) {

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM countries WHERE ID = $1", id)
	if err != nil {
		db.Close()
		return &Country{}, err
	}

	var country Country
	for rows.Next() {
		err := rows.Scan(&country.ID, &country.Name, &country.CreatedAt, &country.UpdatedAt, &country.DeletedAt)
		if err != nil {
			return &Country{}, err
		}
	}

	db.Close()
	return &country, nil
}

func CreateInDB(country *Country) error {
	db := database.DB()
	_, dbErr := db.Query("INSERT INTO countries (name, created_at) VALUES ($1, $2)", country.Name, time.Now())

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}

func UpdateInDB(country *Country) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE countries SET name = $2, updated_at = $3 WHERE id = $1", country.ID, country.Name, time.Now())

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
	_, dbErr := db.Query("UPDATE countries SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	if dbErr != nil {
		db.Close()
		return dbErr
	} else {
		db.Close()
		return nil
	}
}
