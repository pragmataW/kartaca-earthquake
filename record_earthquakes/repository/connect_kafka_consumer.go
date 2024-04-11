package repository

import "github.com/IBM/sarama"

func connectConsumer(broker string) (sarama.Consumer, error){
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil{
		return nil, err
	}

	return conn, nil
}