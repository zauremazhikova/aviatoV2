package passenger

import (
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	responsePassenger, err := GetAllFromDB()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passengers not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Passengers Found", "data": responsePassenger})
}

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")

	responsePassenger, err := GetSingleFromDB(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Passenger Found", "data": responsePassenger})
}

func Create(c *fiber.Ctx) error {
	passenger := new(Passenger)
	err := c.BodyParser(passenger)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = CreateInDB(passenger)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Passenger", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Passenger has created"})
}

func Update(c *fiber.Ctx) error {
	type updatePassenger struct {
		Name string `json:"name"`
	}
	id := c.Params("id")
	passenger, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": err})
	}

	var updatePassengerData updatePassenger
	err = c.BodyParser(&updatePassengerData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	passenger.Name = updatePassengerData.Name
	err = UpdateInDB(passenger)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Passenger has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Passenger has Updated", "data": passenger})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	passenger, err := GetSingleFromDB(id)

	if passenger.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Passenger not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Passenger", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Passenger deleted"})
}
