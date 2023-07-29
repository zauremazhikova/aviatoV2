package entities

import (
	"aviatoV2/database"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

type Passenger struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Passport  string    `json:"passport"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func CreateResponsePassenger(id int, name string, passport string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) Passenger {
	return Passenger{
		ID:        id,
		Name:      name,
		Passport:  passport,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func GetPassengerByID(ID int) Passenger {

	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, PASSPORT, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM passengers WHERE ID = $1", ID)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id        int
		name      string
		passport  string
		createdAt time.Time
		updatedAt time.Time
		deletedAt time.Time
	)

	var passenger Passenger

	for rows.Next() {
		err := rows.Scan(&id, &name, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Fatal(err)
		}
		passenger = CreateResponsePassenger(id, name, passport, createdAt, updatedAt, deletedAt)

	}
	return passenger
}

/*


func GetPassengerByID(id int) Passenger {
	var passengerModel models.Passenger
	FindPassenger(id, &passengerModel)
	passenger := CreateResponsePassenger(passengerModel)
	return passenger
}

func GetPassengers() []Passenger {
	var passengersModel []models.Passenger
	var passengers []Passenger
	FindAllPassengers(&passengersModel)
	for _, passengerModel := range passengersModel {
		passenger := CreateResponsePassenger(passengerModel)
		passengers = append(passengers, passenger)
	}
	return passengers
}

// FindAllPassengers from DB
func FindAllPassengers(country *[]models.Passenger) {
	database.DB.Db.Find(&country)
}

// FindPassenger by ID
func FindPassenger(id int, passenger *models.Passenger) {
	database.DB.Db.Find(&passenger, "id = ?", id)
}

// CreatePassenger in DB
func CreatePassenger(c *fiber.Ctx) error {
	db := database.DB.Db
	passenger := new(models.Passenger)

	err := c.BodyParser(passenger)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&passenger).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Passenger", "data": err})
	}

	responsePassenger := CreateResponsePassenger(*passenger)
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Passenger has created", "data": responsePassenger})
}

// GetAllPassengers from db
func GetAllPassengers(c *fiber.Ctx) error {
	db := database.DB.Db
	var passengers []models.Passenger

	db.Find(&passengers)

	if len(passengers) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passengers not found", "data": nil})
	}

	var responsePassengers []Passenger
	for _, passenger := range passengers {
		responsePassenger := CreateResponsePassenger(passenger)
		responsePassengers = append(responsePassengers, responsePassenger)
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Passengers Found", "data": responsePassengers})
}

// Working with DB

// GetSinglePassenger from db
func GetSinglePassenger(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")
	var passenger models.Passenger

	db.Find(&passenger, "id = ?", id)
	if passenger.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": nil})
	}

	responsePassenger := CreateResponsePassenger(passenger)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Passenger Found", "data": responsePassenger})
}

// UpdatePassenger in db
func UpdatePassenger(c *fiber.Ctx) error {
	type updatePassenger struct {
		Name     string `json:"name"`
		Passport string `json:"passport"`
	}
	db := database.DB.Db
	var passenger models.Passenger

	id := c.Params("id")

	db.Find(&passenger, "id = ?", id)
	if passenger.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": nil})
	}
	var updatePassengerData updatePassenger
	err := c.BodyParser(&updatePassengerData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	passenger.Name = updatePassengerData.Name
	passenger.Passport = updatePassengerData.Passport
	db.Save(&passenger)

	responsePassenger := CreateResponsePassenger(passenger)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Passenger Found", "data": responsePassenger})
}

// DeletePassengerByID in db
func DeletePassengerByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var passenger models.Passenger

	id := c.Params("id")

	db.Find(&passenger, "id = ?", id)
	if passenger.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": nil})
	}
	err := db.Delete(&passenger, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete passenger", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Passenger deleted"})
}
*/
