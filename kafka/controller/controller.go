package controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pragmataW/kartaca-earthquake/kafka/models"
)

func (k *KafkaController) KafkaEarthquakeController(c *fiber.Ctx) error{
	var reqBody kafkaRequest

	if err := c.BodyParser(&reqBody); err != nil{
		fmt.Println(reqBody)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"response": "bad request",
			"error": "wrong format",
		})
	}

	if err := k.IKafkaService.SendMessage(reqBody.Message, reqBody.BrokerAddr, reqBody.Topic, reqBody.Partition); err != nil{
		if _, ok := err.(*models.CouldNotConnectedToBrokerError); ok{
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"response": "unprocessable entity",
				"error": "incorrect broker address",
			})
		}else if _, ok := err.(*models.MessageCouldNotSentToKafkaError); ok{
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"response": "unprocessable entity",
				"error": "incorrect topic name",
			})
		}else if _, ok := err.(*models.CouldNotFindPartitionError); ok{
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"response": "unprocessable entity",
				"error": "incorrect partition id",
			})
		}else{
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"response": "internal server error",
				"error": "undefined error",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"response": "status ok",
	})
}