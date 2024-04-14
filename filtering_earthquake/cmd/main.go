package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/pragmataW/kartaca-earthquake/filtering_earthquake/controller"
	"github.com/pragmataW/kartaca-earthquake/filtering_earthquake/service"
	"github.com/rs/cors"
)

var (
	KafkaServer string
	Topic       string
	Partition   int
)

func main() {
	s := sse.NewServer(nil)
	defer s.Shutdown()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:120"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type"},
	})
	handler := c.Handler(s)

	srv := service.NewFilteringService(KafkaServer, Topic, Partition)
	cntrl := controller.NewFilteringController(srv, s)

	
	go func() {
		if err := cntrl.Srv.HashEarthquakes(); err != nil {
			log.Fatal(err)
		}
		}()
	log.Println("Old keys loading...")
	time.Sleep(5 * time.Second)
	
	http.Handle("/events/", handler)
	
	cntrl.HandleEvents()
	http.HandleFunc("/getOldKeys", cntrl.GetOldKeys)

	http.ListenAndServe(":6663", nil)
}

func init() {
	KafkaServer = os.Getenv("BROKER_SERVER")
	Topic = os.Getenv("TOPIC")

	var err error
	partitionStr := os.Getenv("PARTITION")
	Partition, err = strconv.Atoi(partitionStr)
	if err != nil {
		log.Fatal(err)
	}
}
