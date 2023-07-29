package airline

import (
	"github.com/gofiber/fiber/v2"
)

func GetSingleAirline(c *fiber.Ctx) error {
	id := c.Params("id")

	responseAirline, err := FindOne(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Airline not found", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Airline Found", "data": responseAirline})
}
