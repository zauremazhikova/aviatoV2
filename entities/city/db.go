package city

import (
	"aviatoV2/database"
	"aviatoV2/entities/country"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func GetAllFromDB() (a []*City, err error) {
	cities := make([]*City, 0)

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, COUNTRY_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM cities")
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	for rows.Next() {
		var city City
		var countryID string
		err := rows.Scan(&city.ID, &city.Name, &countryID, &city.CreatedAt, &city.UpdatedAt, &city.DeletedAt)
		if err != nil {
			fmt.Println(err)
			return cities, err
		} else {
			currentCountry, _ := country.GetSingleFromDB(countryID)
			city.Country = *currentCountry
			cities = append(cities, &city)
		}
	}

	return cities, nil
}

func GetSingleFromDB(id string) (*City, error) {

	db := database.DB()
	rows, dbErr := db.Query("SELECT ID, NAME, COUNTRY_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM cities WHERE ID = $1", id)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	var city City
	var countryID string
	for rows.Next() {
		err := rows.Scan(&city.ID, &city.Name, &countryID, &city.CreatedAt, &city.UpdatedAt, &city.DeletedAt)
		if err != nil {
			return &City{}, err
		}
	}
	currentCountry, _ := country.GetSingleFromDB(countryID)
	city.Country = *currentCountry
	return &city, nil

}

func CreateInDB(city *City) error {
	db := database.DB()
	_, dbErr := db.Query("INSERT INTO cities (name, country_id, created_at) VALUES ($1, $2, $3)", city.Name, city.Country.ID, time.Now())

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func UpdateInDB(city *City) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE cities SET name = $2, updated_at = $3 WHERE id = $1", city.ID, city.Name, time.Now())

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, dbErr := db.Query("UPDATE cities SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	if dbErr != nil {
		return dbErr
	} else {
		return nil
	}
}
