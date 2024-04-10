package models

type IncorrectEarthquakeRangeError struct{
	Message string
}

func (e *IncorrectEarthquakeRangeError) Error() string{
	return e.Message
}