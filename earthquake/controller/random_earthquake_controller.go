package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (e EarthquakeController) StartRandomEarthquake(c *fiber.Ctx) error {
	id := e.EarthquakeService.CreateRandomEarthquake(e.KafkaEndpoint, e.BrokerAddr, e.Topic, 0)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func (e EarthquakeController) StopRandomEarthquakeService(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect id parameter",
			"error":   "bad request",
		})
	}

	if err := e.EarthquakeService.StopRandomEarthquake(id); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "earthquake creator id not found",
			"error":   "unprocessable entitiy",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "earthquake service stopped id: " + id,
	})
}
