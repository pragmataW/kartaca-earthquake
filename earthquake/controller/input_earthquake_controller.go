package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pragmataW/kartaca-earthquake/earthquake/models"
)

func (e EarthquakeController) InputEarthquakeController(c *fiber.Ctx) error {
	var body models.Earthquake

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect request format",
			"error":   "bad request",
		})
	}

	earthquake, err := e.EarthquakeService.CreateInputEarthquake(
		body.Lat, body.Lon, body.Mag, e.BrokerAddr, e.Topic, 0,
	)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "invalid earthquake range",
			"error":   "unprocessable entity",
		})
	}

	jsonEarthquake, err := json.Marshal(earthquake)
	if err != nil{
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", e.KafkaEndpoint, bytes.NewBuffer(jsonEarthquake))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "can not reach to kafka",
			"error":   "internal server error",
		})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "can not reach to kafka",
			"error":   "internal server error",
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect kafka request",
			"error":   "bad request",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "earthquake created",
	})
}
