package service

import "github.com/pragmataW/kartaca-earthquake/record_earthquakes/models"

type IRecordEarthquakeRepo interface {
	InsertEarthquake(models.Earthquake) error
}

type RecordEarthquakeService struct {
	Repo   IRecordEarthquakeRepo
	Broker string
	Topic  string
}
