package service

import (
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/repo"
)

func NewService(repository *repo.Repo, broker string, topic string) RecordEarthquakeService {
	return RecordEarthquakeService{
		Repo: repository,
		Broker: broker,
		Topic: topic,
	}
}

func (s RecordEarthquakeService) InsertEarthquakeFromKafka() error {
	worker, err := connectConsumer(s.Broker)
	if err != nil {
		return err
	}
	consumer, err := worker.ConsumePartition(s.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}
	defer consumer.Close()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				parsedData, err := parseData(string(msg.Value))
				if err != nil {
					fmt.Println(err)
				} else {
					err := s.Repo.InsertEarthquake(parsedData)
					if err != nil {
						fmt.Println(err)
					} else {
						log.Println("message accepted")
					}
				}
			}
		}
	}()

	wg.Wait()
	return nil
}