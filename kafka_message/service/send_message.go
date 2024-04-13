package service

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/pragmataW/kafka_message/models"
)

func (k KafkaService) SendMessage(message string, broker string, topic string, partition int32) error {
	client, err := sarama.NewClient([]string{broker}, sarama.NewConfig())
	if err != nil{
		return err
	}
	defer client.Close()

	partitions, err := client.Partitions(topic)
	if err != nil{
		return err
	}

	isThere := false
	for _, p := range partitions{
		if p == partition{
			isThere = true
			break
		}
	}

	if !isThere {
		return &models.CouldNotFindPartitionError{
			Message: fmt.Sprintf("there is no partition id - %d", partition),
		}
	}

	err = PushMessageToQueue(broker, message, topic, partition)
	if err != nil{
		return err
	}
	return nil
}