package country

import (
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	responseCountry, err := GetAllFromDB()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Countries not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Countries Found", "data": responseCountry})
}

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")

	responseCountry, err := GetSingleFromDB(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Country Found", "data": responseCountry})
}

func Create(c *fiber.Ctx) error {
	country := new(Country)
	err := c.BodyParser(country)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = CreateInDB(country)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Country", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Country has created"})
}

func Update(c *fiber.Ctx) error {
	type updateCountry struct {
		Name string `json:"name"`
	}
	id := c.Params("id")
	country, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": err})
	}

	var updateCountryData updateCountry
	err = c.BodyParser(&updateCountryData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	country.Name = updateCountryData.Name
	err = UpdateInDB(country)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Country has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Country has Updated", "data": country})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	country, err := GetSingleFromDB(id)

	if country.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Country not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Country", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Country deleted"})
}
