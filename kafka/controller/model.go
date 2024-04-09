package controller

type kafkaRequest struct {
	Message		string	`json:"message"`
	BrokerAddr	string 	`json:"brokerAddr"`
	Topic		string 	`json:"topic"`
	Partition	int32	`json:"partition"`
}

type IKafkaService interface {
	SendMessage(message string, broker string, topic string, partition int32) error 
}

type KafkaController struct {
	IKafkaService
}

func NewKafkaController(i IKafkaService) *KafkaController {
	ret := new(KafkaController)
	ret.IKafkaService = i

	return ret
}
