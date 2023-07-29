package entities

import (
	"time"
)

type Airline struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func CreateResponseAirline(id int, name string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) Airline {
	return Airline{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

/*

func GetAirlineByID(airlineID int) Airline {

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM airlines WHERE ID = $1", airlineID)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id        int
		name      string
		createdAt time.Time
		updatedAt time.Time
		deletedAt time.Time
	)

	var airline Airline

	for rows.Next() {
		err := rows.Scan(&id, &name, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Fatal(err)
		}
		airline = CreateResponseAirline(id, name, createdAt, updatedAt, deletedAt)

	}
	return airline
}







func GetAirlines() []Airline {
	var airlinesModel []Airline
	var airlines []Airline
	FindAllAirlines(&airlinesModel)
	for _, airlineModel := range airlinesModel {
		airline := CreateResponseAirline(airlineModel)
		airlines = append(airlines, airline)
	}
	return airlines
}

// FindAllAirlines from DB
func FindAllAirlines(airline *[]Airline) {
	database.DB.Db.Find(&airline)
}

// FindAirline by ID
func FindAirline(id int, airline *Airline) {
	database.DB.Db.Find(&airline, "id = ?", id)
}

// Working with DB

// CreateAirline in DB
func CreateAirline(c *fiber.Ctx) error {
	db := database.DB.Db
	airline := new(Airline)
	err := c.BodyParser(airline)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&airline).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Airline", "data": err})
	}

	responseAirline := CreateResponseAirline(*airline)
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline has created", "data": responseAirline})
}

// GetAllAirlines from db
func GetAllAirlines(c *fiber.Ctx) error {
	db := database.DB.Db
	var airlines []Airline
	db.Find(&airlines)
	if len(airlines) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airlines not found", "data": nil})
	}

	var responseAirlines []Airline
	for _, airline := range airlines {
		responseAirline := CreateResponseAirline(airline)
		responseAirlines = append(responseAirlines, responseAirline)
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Airlines Found", "data": responseAirlines})
}

// GetSingleAirline from db
func GetSingleAirline(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var airline Airline

	db.Find(&airline, "id = ?", id)
	if airline.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": nil})
	}
	responseAirline := CreateResponseAirline(airline)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Airline Found", "data": responseAirline})
}

// UpdateAirline in db
func UpdateAirline(c *fiber.Ctx) error {
	type updateAirline struct {
		Name string `json:"name"`
	}
	db := database.DB.Db

	var airline Airline
	id := c.Params("id")
	db.Find(&airline, "id = ?", id)
	if airline.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": nil})
	}

	var updateAirlineData updateAirline
	err := c.BodyParser(&updateAirlineData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	airline.Name = updateAirlineData.Name
	db.Save(&airline)
	responseAirline := CreateResponseAirline(airline)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "airline Found", "data": responseAirline})
}

// DeleteAirlineByID in db
func DeleteAirlineByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var airline Airline
	id := c.Params("id")
	db.Find(&airline, "id = ?", id)

	if airline.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": nil})
	}

	err := db.Delete(&airline, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete Airline", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Airline deleted"})
}

*/
