package direction

import (
	"aviatoV2/entities/airline"
	"aviatoV2/entities/city"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	responseDirection, err := GetAllFromDB()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Directions not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Directions Found", "data": responseDirection})
}

func GetSingle(c *fiber.Ctx) error {
	id := c.Params("id")

	responseDirection, err := GetSingleFromDB(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Direction not found", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Direction Found", "data": responseDirection})
}

func Create(c *fiber.Ctx) error {
	type insertStruct struct {
		OriginCityID      string `json:"originCityID"`
		DestinationCityID string `json:"destinationCityID"`
		AirlineID         string `json:"airlineID"`
	}
	var insertData insertStruct
	err := c.BodyParser(&insertData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	direction := new(Direction)
	originCity, _ := city.GetSingleFromDB(insertData.OriginCityID)
	destinationCity, _ := city.GetSingleFromDB(insertData.DestinationCityID)
	airlineCity, _ := airline.GetSingleFromDB(insertData.AirlineID)

	direction.OriginCity = *originCity
	direction.DestinationCity = *destinationCity
	direction.Airline = *airlineCity

	err = CreateInDB(direction)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Direction", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Direction has created"})
}

func Update(c *fiber.Ctx) error {
	type updateStruct struct {
		OriginCityID      string `json:"originCityID"`
		DestinationCityID string `json:"destinationCityID"`
		AirlineID         string `json:"airlineID"`
	}
	id := c.Params("id")
	direction, err := GetSingleFromDB(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Direction not found", "data": err})
	}

	var updateData updateStruct
	err = c.BodyParser(&updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	originCity, _ := city.GetSingleFromDB(updateData.OriginCityID)
	destinationCity, _ := city.GetSingleFromDB(updateData.DestinationCityID)
	airlineCity, _ := airline.GetSingleFromDB(updateData.AirlineID)

	direction.OriginCity = *originCity
	direction.DestinationCity = *destinationCity
	direction.Airline = *airlineCity

	err = UpdateInDB(direction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Direction has not updated", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Direction has Updated", "data": direction})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	direction, err := GetSingleFromDB(id)

	if direction.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Direction not found", "data": nil})
	}

	err = DeleteInDB(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete Direction", "data": nil})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Direction deleted"})
}
