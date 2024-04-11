package controller

import "github.com/pragmataW/kartaca-earthquake/record_earthquakes/models"

type IRecordEarthquakeService interface {
	InsertEarthquakeFromKafka() error
	SelectEarthquakeFromSql() ([]models.Earthquake, error)
}

type RecordEarthquakeController struct{
	Services IRecordEarthquakeService
}
