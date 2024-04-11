package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewController(srv IRecordEarthquakeService) RecordEarthquakeController {
	return RecordEarthquakeController{
		Services: srv,
	}
}

func (ct RecordEarthquakeController) GetEarthquakeDatas(c *fiber.Ctx) error {
	earthquakes, err :=  ct.Services.SelectEarthquakeFromSql()
	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"error": "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(earthquakes)
}
