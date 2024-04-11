package service

type IRecordEarthquakeRepo interface{
	InsertEarthquakeFromKafka() error
}

type RecordEarthquakeService struct{
	Repo IRecordEarthquakeRepo
}