package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/controller"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/models"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/repo"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/service"
)

var (
	Host        string
	Port        int
	User        string
	Password    string
	DbName      string
	SslMode     string
	KafkaServer string
	Topic       string
)

func main() {
	repo := repo.NewRepository(models.RepoConfig{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		DbName:   DbName,
		SslMode:  SslMode,
		Broker:   KafkaServer,
		Topic:    Topic,
	})
	srv := service.NewService(repo, KafkaServer, Topic)
	ctrlr := controller.NewController(srv)

	go func() {
		if err := srv.InsertEarthquakeFromKafka(); err != nil {
			log.Fatal(err)
		}
	}()

	app := fiber.New()
	app.Get("/getEarthquakes", ctrlr.GetEarthquakeDatas)
	app.Listen(":3232")
}

func init() {
	Host = os.Getenv("HOST")
	User = os.Getenv("DB_USERNAME")
	Password = os.Getenv("PASSWORD")
	DbName = os.Getenv("DB_NAME")
	SslMode = os.Getenv("SSL_MODE")
	KafkaServer = os.Getenv("KAFKA_SERVER")
	Topic = os.Getenv("TOPIC")

	portStr := os.Getenv("PORT")
	Port, _ = strconv.Atoi(portStr)
}
