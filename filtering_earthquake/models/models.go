package models

import (
	"log"
	"sync"
)

type Earthquake struct {
	Lat float64
	Lon float64
	Mag float64
}

type ObservableMap struct {
	Earthquakes map[string][]Earthquake
	Observers   []func(string, Earthquake)
	mu          sync.Mutex
}

func NewObservableMap() *ObservableMap {
	return &ObservableMap{
		Earthquakes: make(map[string][]Earthquake),
		Observers:   make([]func(string, Earthquake), 0),
	}
}

func (o *ObservableMap) Set(key string, value Earthquake) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if value.Mag < 3{
		return
	}

	log.Println("Set", key, value)
	o.Earthquakes[key] = append(o.Earthquakes[key], value)
	o.notifyObservers(key, value)
}

func (o *ObservableMap) notifyObservers(key string, value Earthquake) {
	for _, observer := range o.Observers {
		observer(key, value)
	}
}

func (o *ObservableMap) RegisterObserver(observer func(string, Earthquake)) {
	o.Observers = append(o.Observers, observer)
}

var GroupedEarthquakes = NewObservableMap()
