package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/pragmataW/kartaca-earthquake/filtering_earthquake/models"
)

func NewFilteringController(srv IFilteringService, sv *sse.Server) FilteringController {
	return FilteringController{
		Srv:   srv,
		SseSv: sv,
	}
}

func (c FilteringController) HandleEvents() {
	for keys := range models.GroupedEarthquakes.Earthquakes {
		OldKeys[keys] = len(models.GroupedEarthquakes.Earthquakes[keys])
	}

	models.GroupedEarthquakes.RegisterObserver(func(key string, value models.Earthquake) {
		len := len(models.GroupedEarthquakes.Earthquakes[key])
		c.SseSv.SendMessage("/events/message", sse.SimpleMessage(fmt.Sprintf("%s,%d", key, len)))
	})

	models.GroupedEarthquakes.RegisterObserver(func(key string, value models.Earthquake) {
		mutex.Lock()
		defer mutex.Unlock()
		OldKeys[key] = len(models.GroupedEarthquakes.Earthquakes[key])
		log.Println("old keys added")
	})
}

func (c FilteringController) GetOldKeys(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	for key, value := range OldKeys {
		fmt.Fprintf(w, "%s,%d\n", key, value)
	}
}
