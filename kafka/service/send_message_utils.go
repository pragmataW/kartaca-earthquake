package service

import (
	"github.com/IBM/sarama"
	"github.com/pragmataW/kartaca-earthquake/kafka/models"
)

//! Producer

func ConnectProducer(broker string) (sarama.SyncProducer, error){
	brokerAddr := []string{broker} 
	
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokerAddr, config)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func PushMessageToQueue(broker string, messageStr string, topic string, partition int32) error{
	producer, err := ConnectProducer(broker)
	if err != nil{
		return &models.CouldNotConnectedToBrokerError{
			Message: "could not connected to broker",
		}
	}
	
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(messageStr),
		Partition: partition,
	}

	_, _, err = producer.SendMessage(message)
	if err != nil{
		return &models.MessageCouldNotSentToKafkaError{
			Message: "message could not sent to kafka",
		}
	}
	return nil
}