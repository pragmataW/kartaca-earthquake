package service

import (
	"log"

	"github.com/IBM/sarama"
)

func NewFilteringService(brokerAddr string, topic string, partition int) FilteringService {
	return FilteringService{
		KafkaServer: brokerAddr,
		Topic:       topic,
		Partition: partition,
	}
}

func (s FilteringService) HashEarthquakes() error {
	worker, err := connectConsumer(s.KafkaServer)
	if err != nil {
		return err
	}
	defer worker.Close()

	consumer, err := worker.ConsumePartition(s.Topic, int32(s.Partition), sarama.OffsetOldest)
	if err != nil{
		return err
	}
	defer consumer.Close()

	for {
		select {
		case err := <-consumer.Errors():
			return err
		case msg := <-consumer.Messages():
			log.Println(string(msg.Value))
			err := processMessages(string(msg.Value))
			if err != nil{
				return err
			}
		}
	}
}
