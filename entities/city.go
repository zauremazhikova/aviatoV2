package entities

import (
	"time"
)

type City struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Country   Country   `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

/*
func CreateResponseCity(id int, name string, countryID int, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) City {

	return City{
		ID:        id,
		Name:      name,
		Country:   GetCountryByID(countryID),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func GetCityByID(ID int) City {
	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, COUNTRY_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM cities WHERE ID = $1", ID)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id        int
		name      string
		countryID int
		createdAt time.Time
		updatedAt time.Time
		deletedAt time.Time
	)

	var city City
	for rows.Next() {
		err := rows.Scan(&id, &name, &countryID, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Fatal(err)
		}
		city = CreateResponseCity(id, name, countryID, createdAt, updatedAt, deletedAt)

	}
	return city
}




func GetCityByID(id int) City {
	var cityModel models.City
	FindCity(id, &cityModel)
	city := CreateResponseCity(cityModel)
	return city
}

func GetCities() []City {
	var citiesModel []models.City
	var cities []City
	FindAllCities(&citiesModel)
	for _, cityModel := range citiesModel {
		city := CreateResponseCity(cityModel)
		cities = append(cities, city)
	}
	return cities
}

// FindAllCities from DB
func FindAllCities(city *[]models.City) {
	database.DB.Db.Find(&city)
}

// FindCity by ID
func FindCity(id int, city *models.City) {
	database.DB.Db.Find(&city, "id = ?", id)
}

// Working with DB

// CreateCity in DB
func CreateCity(c *fiber.Ctx) error {
	db := database.DB.Db
	city := new(models.City)
	err := c.BodyParser(city)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&city).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create City", "data": err})
	}

	responseCity := CreateResponseCity(*city)
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City has created", "data": responseCity})
}

// GetAllCities from db
func GetAllCities(c *fiber.Ctx) error {
	db := database.DB.Db
	var cities []models.City
	db.Find(&cities)
	if len(cities) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cities not found", "data": nil})
	}

	var responseCities []City
	for _, city := range cities {
		responseCity := CreateResponseCity(city)
		responseCities = append(responseCities, responseCity)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Cities Found", "data": responseCities})
}

// GetSingleCity from db
func GetSingleCity(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")
	var city models.City

	db.Find(&city, "id = ?", id)
	if city.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": nil})
	}
	responseCity := CreateResponseCity(city)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "City Found", "data": responseCity})
}

// UpdateCity in db
func UpdateCity(c *fiber.Ctx) error {
	type updateCity struct {
		Name      string `json:"name"`
		CountryID int    `json:"CountryID"`
	}
	db := database.DB.Db
	var city models.City

	id := c.Params("id")

	db.Find(&city, "id = ?", id)
	if city.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": nil})
	}
	var updateCityData updateCity
	err := c.BodyParser(&updateCityData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	city.Name = updateCityData.Name
	city.CountryID = updateCityData.CountryID
	db.Save(&city)
	responseCity := CreateResponseCity(city)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "city Found", "data": responseCity})
}

// DeleteCityByID in db
func DeleteCityByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var city models.City
	// get id params
	id := c.Params("id")
	// find single city in the database by id
	db.Find(&city, "id = ?", id)
	if city.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": nil})
	}
	err := db.Delete(&city, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete city", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "City deleted"})
}
*/
