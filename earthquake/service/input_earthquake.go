package service

import (
	"fmt"

	"github.com/pragmataW/kartaca-earthquake/earthquake/models"
)

func (e EarthquakeService) CreateInputEarthquake(lat float64, lon float64, magnitude float64, brokerAddr string, topic string, partition int32) (models.KafkaRequest, error) {
	earthquake := models.Earthquake{
		Lat: lat,
		Lon: lon,
		Mag: magnitude,
	}

	if !earthquake.ControlDatas(){
		return models.KafkaRequest{}, &models.IncorrectEarthquakeRangeError{
			Message: "incorrect earthquake value error",
		}
	}
	
	ret := models.KafkaRequest{
		Message: fmt.Sprintf("lat:%f,lon:%f,mag:%f", earthquake.Lat, earthquake.Lon, earthquake.Mag),
		BrokerAddr: brokerAddr,
		Topic: topic,
		Partition: partition,
	}
	return ret, nil
}