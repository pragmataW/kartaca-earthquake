package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/models"
	_ "github.com/lib/pq"
)

func NewRepository(cfg models.RepoConfig) *Repo {
	fmt.Println(cfg)
	once.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		if err := db.Ping(); err != nil{
			log.Fatal(err)
		}

		repo = &Repo{Db: db}
		repo.Broker = cfg.Broker
		repo.Topic = cfg.Topic
		fmt.Println("repo obj created")
	})
	return repo
}

func (r Repo) InsertEarthquakeFromKafka() error {
	worker, err := connectConsumer(r.Broker)
	if err != nil {
		return err
	}

	consumer, err := worker.ConsumePartition(r.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}
	defer consumer.Close()

	stmt, err := r.Db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				parsedData, err := parseData(string(msg.Value))
				if err != nil {
					fmt.Println(err)
				} else {
					_, err := stmt.Exec(parsedData.Lat, parsedData.Lon, parsedData.Mag)
					if err != nil {
						fmt.Println(err)
					}else{
						log.Println("message accepted")
					}
				}
			}
		}
	}()

	wg.Wait()
	return nil
}


//! REFACTORING YAP, SADECE STMT.EXEC KISIMLARINI REPOSITORY'DE BIRAK, GERİSİNİ SERVİCE KISMINA TAŞI