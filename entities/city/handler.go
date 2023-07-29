package city

import "github.com/gofiber/fiber/v2"

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
	city := new(City)
	err := c.BodyParser(city)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = CreateInDB(city)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create City", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City has created"})
}

func Update(c *fiber.Ctx) error {
	type updateCity struct {
		Name string `json:"name"`
	}
	id := c.Params("id")
	city, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": err})
	}

	var updateCityData updateCity
	err = c.BodyParser(&updateCityData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	city.Name = updateCityData.Name
	err = UpdateInDB(city)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "City has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City has Updated", "data": city})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	city, err := GetSingleFromDB(id)

	if city.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "City not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete City", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "City deleted"})
}
