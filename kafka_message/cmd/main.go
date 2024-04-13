package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	kafkacontroller "github.com/pragmataW/kafka_message/controller"
	"github.com/pragmataW/kafka_message/service"
)

func main() {
	var kafkaService service.KafkaService

	KafkaController := kafkacontroller.KafkaController{
		IKafkaService: kafkaService,
	}

	app := fiber.New()
	app.Post("/sendMessageToKafka", KafkaController.KafkaEarthquakeController)

	log.Fatal(app.Listen(":3030"))
}