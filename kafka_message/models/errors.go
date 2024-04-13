package models

type CouldNotConnectedToBrokerError struct{
	Message string
}

type MessageCouldNotSentToKafkaError struct{
	Message string
}

type CouldNotFindPartitionError struct{
	Message string
}

func (e *CouldNotConnectedToBrokerError) Error() string{
	return e.Message
}

func (e *MessageCouldNotSentToKafkaError) Error() string {
	return e.Message
}

func (e *CouldNotFindPartitionError) Error() string {
	return e.Message
}