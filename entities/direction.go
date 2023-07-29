package entities

import (
	"aviatoV2/entities/airline"
	"time"
)

type Direction struct {
	ID              int             `json:"id"`
	OriginCity      City            `json:"originCity"`
	DestinationCity City            `json:"destinationCity"`
	Airline         airline.Airline `json:"airline"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `json:"deleted_at"`
}

/*
func CreateResponseDirection(id int, originCityID int, destinationCityID int, airlineID int, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) Direction {
	return Direction{
		ID:              id,
		OriginCity:      GetCityByID(originCityID),
		DestinationCity: GetCityByID(destinationCityID),
		Airline:         GetAirlineByID(airlineID),
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		DeletedAt:       deletedAt,
	}
}

func GetDirectionByID(ID int) Direction {
	db := database.DB()
	rows, err := db.Query("SELECT ID, ORIGIN_CITY_ID, DESTINATION_CITY_ID, AIRLINE_ID, CREATED_AT, COALESCE(UPDATED_AT, DATE('0001-01-01')) AS UPDATED_AT, COALESCE(DELETED_AT, DATE('0001-01-01')) AS DELETED_AT FROM destination WHERE ID = $1", ID)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id                int
		originCityID      int
		destinationCityID int
		airlineCityID     int
		createdAt         time.Time
		updatedAt         time.Time
		deletedAt         time.Time
	)

	var direction Direction
	for rows.Next() {
		err := rows.Scan(&id, &originCityID, &destinationCityID, &airlineCityID, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Fatal(err)
		}
		direction = CreateResponseDirection(id, originCityID, destinationCityID, airlineCityID, createdAt, updatedAt, deletedAt)

	}
	return direction
}


func GetRouteByID(id int) Route {
	var routeModel models.Route
	FindRoute(id, &routeModel)
	route := CreateResponseRoute(routeModel)
	return route
}

func GetRoutes() []Route {
	var routesModel []models.Route
	var routes []Route
	FindAllRoutes(&routesModel)
	for _, routeModel := range routesModel {
		route := CreateResponseRoute(routeModel)
		routes = append(routes, route)
	}
	return routes
}

// GetRoutesByOriginDestination from DB
func GetRoutesByOriginDestination(originID int, destinationID int) []Route {
	var routesModel []models.Route
	var routes []Route
	FindRoutesByOriginDestination(originID, destinationID, &routesModel)
	for _, routeModel := range routesModel {
		route := CreateResponseRoute(routeModel)
		routes = append(routes, route)
	}
	return routes
}

// GetRoutesByOriginCity from DB
func GetRoutesByOriginCity(originID int) []Route {
	var routesModel []models.Route
	var routes []Route
	FindRoutesByOriginCity(originID, &routesModel)
	for _, routeModel := range routesModel {
		route := CreateResponseRoute(routeModel)
		routes = append(routes, route)
	}
	return routes
}

// FindAllRoutes from DB
func FindAllRoutes(routes *[]models.Route) {
	database.DB.Db.Find(&routes)
}

// FindRoute by ID
func FindRoute(id int, route *models.Route) {
	database.DB.Db.Find(&route, "id = ?", id)
}

// FindRoutesByOriginDestination by ID
func FindRoutesByOriginDestination(originID int, destinationID int, route *[]models.Route) {
	database.DB.Db.Where("origin_city_id = ?", originID).Where("destination_city_id = ?", destinationID).Find(&route)
}

// FindRoutesByOriginCity by ID
func FindRoutesByOriginCity(originID int, route *[]models.Route) {
	database.DB.Db.Where("origin_city_id = ?", originID).Find(&route)
}

// Working with DB

// CreateRoute in DB
func CreateRoute(c *fiber.Ctx) error {
	db := database.DB.Db
	route := new(models.Route)

	err := c.BodyParser(route)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&route).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Route", "data": err})
	}

	responseRoute := CreateResponseRoute(*route)

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Route has created", "data": responseRoute})
}

// GetAllRoutes from db
func GetAllRoutes(c *fiber.Ctx) error {
	db := database.DB.Db
	var routes []models.Route

	db.Find(&routes)
	if len(routes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Routes not found", "data": nil})
	}

	var responseRoutes []Route
	for _, route := range routes {
		responseRoute := CreateResponseRoute(route)
		responseRoutes = append(responseRoutes, responseRoute)
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Routes Found", "data": responseRoutes})
}

// GetSingleRoute from db
func GetSingleRoute(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var route models.Route
	// find single route in the database by id
	db.Find(&route, "id = ?", id)
	if route.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Route not found", "data": nil})
	}

	responseRoute := CreateResponseRoute(route)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Route Found", "data": responseRoute})
}

// UpdateRoute in db
func UpdateRoute(c *fiber.Ctx) error {
	type updateRoute struct {
		OriginCityID      int `json:"OriginCityID"`
		DestinationCityID int `json:"DestinationCityID"`
		AirlineID         int `json:"AirlineID"`
	}
	db := database.DB.Db
	var route models.Route

	id := c.Params("id")
	db.Find(&route, "id = ?", id)
	if route.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Route not found", "data": nil})
	}
	var updateRouteData updateRoute
	err := c.BodyParser(&updateRouteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	route.OriginCityID = updateRouteData.OriginCityID
	route.DestinationCityID = updateRouteData.DestinationCityID
	route.AirlineID = updateRouteData.AirlineID
	db.Save(&route)

	responseRoute := CreateResponseRoute(route)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "route Found", "data": responseRoute})
}

// DeleteRouteByID in db
func DeleteRouteByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var route models.Route
	// get id params
	id := c.Params("id")
	// find single route in the database by id
	db.Find(&route, "id = ?", id)
	if route.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Route not found", "data": nil})
	}
	err := db.Delete(&route, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete route", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Route deleted"})
}
*/
