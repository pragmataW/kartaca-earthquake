package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pragmataW/kartaca-earthquake/earthquake/controller"
	"github.com/pragmataW/kartaca-earthquake/earthquake/service"
)

var (
	brokerAddr    string
	topic         string
	kafkaEndpoint string
)

func main() {
	service := service.EarthquakeService{}
	controller := controller.NewEarthquakeController(service, brokerAddr, topic, kafkaEndpoint, 0)

	app := fiber.New()
	app.Post("/inputEarthquake", controller.InputEarthquakeController)
	app.Post("/startRandomEarthquake", controller.StartRandomEarthquake)
	app.Delete("/stopRandomEarthquake/:id", controller.StopRandomEarthquakeService)
	log.Fatal(app.Listen(":3131"))
}

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("env load error")
	}

	brokerAddr = os.Getenv("BROKER_ADDR")
	topic = os.Getenv("TOPIC")
	kafkaEndpoint = os.Getenv("KAFKA_ENDPOINT")
}
