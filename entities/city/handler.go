package city

import (
	"aviatoV2/entities/country"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	responseCity, err := GetAllFromDB()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Cities not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Cities Found", "data": responseCity})
}

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")

	responseCity, err := GetSingleFromDB(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City Found", "data": responseCity})
}

func Create(c *fiber.Ctx) error {
	type insertStruct struct {
		Name      string `json:"name"`
		CountryID string `json:"countryID"`
	}
	var insertData insertStruct
	err := c.BodyParser(&insertData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	currentCountry, _ := country.GetSingleFromDB(insertData.CountryID)
	if currentCountry.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": err})
	}

	city := new(City)
	city.Name = insertData.Name
	city.Country = *currentCountry

	err = CreateInDB(city)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create City", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City has created"})
}

func Update(c *fiber.Ctx) error {
	type updateStruct struct {
		Name      string `json:"name"`
		CountryID string `json:"countryID"`
	}
	id := c.Params("id")
	city, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": err})
	}

	var updateData updateStruct
	err = c.BodyParser(&updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	currentCountry, _ := country.GetSingleFromDB(updateData.CountryID)
	if currentCountry.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": err})
	}

	city.Name = updateData.Name
	city.Country = *currentCountry

	err = UpdateInDB(city)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "City has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City has Updated", "data": city})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	city, err := GetSingleFromDB(id)

	if city.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete City", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City deleted"})
}
