package entities

import (
	"time"
)

type Country struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

/*
func CreateResponseCountry(id int, name string, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) Country {
	return Country{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func GetCountryByID(countryID int) Country {
	db := database.DB()
	rows, err := db.Query("SELECT ID, NAME, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM countries WHERE ID = $1", countryID)
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

	var country Country
	for rows.Next() {
		err := rows.Scan(&id, &name, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Fatal(err)
		}
		country = CreateResponseCountry(id, name, createdAt, updatedAt, deletedAt)

	}
	return country
}


func GetCountryByID(id int) Country {
	var countryModel models.Country
	FindCountry(id, &countryModel)
	country := CreateResponseCountry(countryModel)
	return country
}

func GetCountries() []Country {
	var countriesModel []models.Country
	var countries []Country
	FindAllCountries(&countriesModel)
	for _, countryModel := range countriesModel {
		country := CreateResponseCountry(countryModel)
		countries = append(countries, country)
	}
	return countries
}

// FindAllCountries from DB
func FindAllCountries(country *[]models.Country) {
	database.DB.Db.Find(&country)
}

// FindCountry by ID
func FindCountry(id int, country *models.Country) {
	database.DB.Db.Find(&country, "id = ?", id)
}

// Working with DB

// CreateCountry in DB
func CreateCountry(c *fiber.Ctx) error {
	db := database.DB.Db
	country := new(models.Country)

	err := c.BodyParser(country)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&country).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Country", "data": err})
	}

	responseCountry := CreateResponseCountry(*country)
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Country has created", "data": responseCountry})
}

// GetAllCountries from db
func GetAllCountries(c *fiber.Ctx) error {
	db := database.DB.Db
	var countries []models.Country

	db.Find(&countries)

	if len(countries) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Countries not found", "data": nil})
	}

	var responseCountries []Country
	for _, country := range countries {
		responseCountry := CreateResponseCountry(country)
		responseCountries = append(responseCountries, responseCountry)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Countries Found", "data": responseCountries})
}

// GetSingleCountry from db
func GetSingleCountry(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")
	var country models.Country

	db.Find(&country, "id = ?", id)
	if country.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": nil})
	}

	responseCountry := CreateResponseCountry(country)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Country Found", "data": responseCountry})
}

// UpdateCountry in db
func UpdateCountry(c *fiber.Ctx) error {
	type updateCountry struct {
		Name string `json:"name"`
	}
	db := database.DB.Db
	var country models.Country

	id := c.Params("id")

	db.Find(&country, "id = ?", id)
	if country.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": nil})
	}
	var updateCountryData updateCountry
	err := c.BodyParser(&updateCountryData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	country.Name = updateCountryData.Name
	db.Save(&country)

	responseCountry := CreateResponseCountry(country)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "country Found", "data": responseCountry})
}

// DeleteCountryByID in db
func DeleteCountryByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var country models.Country

	id := c.Params("id")

	db.Find(&country, "id = ?", id)
	if country.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": nil})
	}
	err := db.Delete(&country, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete country", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Country deleted"})
}
*/
