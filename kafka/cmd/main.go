package main

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	kafkacontroller "github.com/pragmataW/kartaca-earthquake/kafka/controller"
	"github.com/pragmataW/kartaca-earthquake/kafka/service"
)

var brokerAddr string

func main() {
	var kafkaService service.KafkaService
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	KafkaController := kafkacontroller.KafkaController{
		IKafkaService: kafkaService,
	}

	app := fiber.New()
	app.Post("/sendMessageToKafka", KafkaController.KafkaEarthquakeController)

	log.Fatal(app.Listen(":3030"))
}

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	brokerAddr = os.Getenv(brokerAddr)
}
