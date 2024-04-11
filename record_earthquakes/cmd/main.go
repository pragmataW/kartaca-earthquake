package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/service"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/models"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/repo"
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
	srv.InsertEarthquakeFromKafka()
}

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

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
