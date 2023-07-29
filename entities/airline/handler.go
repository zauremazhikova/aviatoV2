package airline

import (
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	responseAirline, err := GetAllFromDB()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airlines not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airlines Found", "data": responseAirline})
}

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")

	responseAirline, err := GetSingleFromDB(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline Found", "data": responseAirline})
}

func Create(c *fiber.Ctx) error {
	airline := new(Airline)
	err := c.BodyParser(airline)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = CreateInDB(airline)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Airline", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline has created"})
}

func Update(c *fiber.Ctx) error {
	type updateAirline struct {
		Name string `json:"name"`
	}
	id := c.Params("id")
	airline, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": err})
	}

	var updateAirlineData updateAirline
	err = c.BodyParser(&updateAirlineData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	airline.Name = updateAirlineData.Name
	err = UpdateInDB(airline)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Airline has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline has Updated", "data": airline})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	airline, err := GetSingleFromDB(id)

	if airline.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Airline", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Airline deleted"})
}
