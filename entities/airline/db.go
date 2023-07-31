package airline

import (
	"aviatoV2/database"
	"time"
)

func GetAllFromDB() (a []*Airline, err error) {
	airlines := make([]*Airline, 0)

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines")

	_ = db.Close()
	if err != nil {
		return nil, err
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
	rows, err := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines WHERE ID = $1", id)
	_ = db.Close()
	if err != nil {
		return nil, err
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
	_, err := db.Query("INSERT INTO airlines (name, created_at) VALUES ($1, $2)", airline.Name, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func UpdateInDB(airline *Airline) error {
	db := database.DB()
	_, err := db.Query("UPDATE airlines SET name = $2, updated_at = $3 WHERE id = $1", airline.ID, airline.Name, time.Now())

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeleteInDB(id string) error {
	db := database.DB()
	_, err := db.Query("UPDATE airlines SET deleted_at = $1 WHERE id = $2", time.Now(), id)

	_ = db.Close()
	if err != nil {
		return err
	}
	return nil
}
