package kafka

import (
	"github.com/IBM/sarama"
	commonobj "github.com/pragmataW/kartaca-earthquake/common_obj"
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
		return &commonobj.CouldNotConnectedToBrokerError{
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
		return &commonobj.MessageCouldNotSentToKafkaError{
			Message: "message could not sent to kafka",
		}
	}
	return nil
}

//! Consumer

func ConnectConsumer(broker string) (sarama.Consumer, error){
	brokerAddr := []string{broker}
	
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokerAddr, config)
	if err != nil{
		return nil, err
	}
	return consumer, err
}