package models

import "fmt"

type Earthquake struct {
	Lat 	float64	`json:"lat"`
	Lon 	float64	`json:"lon"`
	Mag 	float64	`json:"mag"`
}

func (e *Earthquake) String() string {
	return fmt.Sprintf("Latitude: %f, Longitude: %f, Magnitude: %f\n", e.Lat, e.Lon, e.Mag)
}

func (e *Earthquake) ControlDatas() bool {
	if (e.Lat > 90 || e.Lat < -90) ||
		(e.Lon > 180 || e.Lon < -180) ||
		(e.Mag > 10 || e.Mag < 1){
			return false
		}
		return true
}
