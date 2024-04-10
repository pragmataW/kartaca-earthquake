package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pragmataW/kartaca-earthquake/earthquake/models"
)

func (e EarthquakeService) CreateRandomEarthquake(kafkaEndpoint string, brokerAddr string, topic string, partition int32) (string) {
	id := uuid.New().String()
	quitCh := make(chan bool)

	job := &models.Job{
		Id:       id,
		QuitChan: quitCh,
	}

	models.JobsMutex.Lock()
	models.Jobs[id] = job
	models.JobsMutex.Unlock()

	go func() {
		for {
			select {
			case <-job.QuitChan:
				return
			default:
				job.Earthquake.Lat = rand.Float64()*180 - 90
				job.Lon = rand.Float64()*360 - 180
				job.Mag = rand.Float64() * 10

				reqDraft := models.KafkaRequest{
					Message:    fmt.Sprintf("lat:%f,lon:%f,mag:%f", job.Lat, job.Lon, job.Mag),
					BrokerAddr: brokerAddr,
					Topic:      topic,
					Partition:  partition,
				}
				jsonReq, _ := json.Marshal(reqDraft)
				req, _ := http.NewRequest("POST", kafkaEndpoint, bytes.NewBuffer(jsonReq))
				req.Header.Set("Content-Type", "application/json")
				client := &http.Client{}
				resp, _ := client.Do(req)
				defer resp.Body.Close()
				log.Println(resp.Status)

				time.Sleep(2 * time.Second)
			}
		}
	}()

	return id
}

func (e EarthquakeService) StopRandomEarthquake(id string) error {
	models.JobsMutex.Lock()
	job, exists := models.Jobs[id]
	if exists{
		close(job.QuitChan)
		delete(models.Jobs, id)
	}
	models.JobsMutex.Unlock()

	if !exists{
		return errors.New("earthquake creator id not found")
	}
	return nil
}