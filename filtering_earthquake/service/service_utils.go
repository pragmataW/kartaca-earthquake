package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IBM/sarama"
	"github.com/mmcloughlin/geohash"
	"github.com/pragmataW/kartaca-earthquake/filtering_earthquake/models"
)

func connectConsumer(broker string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func parseData(data string) (models.Earthquake, error) {
	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return models.Earthquake{}, fmt.Errorf("beklenen veri formatı bulunamadı: %s", data)
	}

	var values []float64
	for _, part := range parts {
		keyValue := strings.Split(part, ":")
		if len(keyValue) != 2 {
			return models.Earthquake{}, fmt.Errorf("beklenen anahtar:değer formatı bulunamadı: %s", part)
		}

		value, err := strconv.ParseFloat(keyValue[1], 64)
		if err != nil {
			return models.Earthquake{}, fmt.Errorf("float'a çevirme hatası: %v", err)
		}

		values = append(values, value)
	}

	return models.Earthquake{
		Lat: values[0],
		Lon: values[1],
		Mag: values[2],
	}, nil
}

func processMessages(data string) error {
	eq, err := parseData(data)
	if err != nil{
		return err
	}

	hash := geohash.EncodeWithPrecision(eq.Lat, eq.Lon, 4)
	models.GroupedEarthquakes.Set(hash, eq)
	return nil
}