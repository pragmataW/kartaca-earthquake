package controller

import "github.com/pragmataW/kartaca-earthquake/earthquake/models"

type IEarthquakeService interface {
	//CreateRandomEarthquake() error
	CreateInputEarthquake(lat float64, lon float64, magnitude float64, brokerAddr string, topic string, partition int32) (models.KafkaRequest, error)
}

type EarthquakeController struct {
	EarthquakeService IEarthquakeService
	BrokerAddr        string
	Topic             string
	KafkaEndpoint     string
}


func NewEarthquakeController(i IEarthquakeService, brokerAddr string, topic string, kafkaEndpoint string, partition int) EarthquakeController {
	return EarthquakeController{
		EarthquakeService: i,
		BrokerAddr: brokerAddr,
		Topic: topic,
		KafkaEndpoint: kafkaEndpoint,
	}
}
