package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pragmataW/kartaca-earthquake/record_earthquakes/models"
)

func NewRepository(cfg models.RepoConfig) *Repo {
	once.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		repo = &Repo{Db: db}
		fmt.Println("repo obj created")
	})
	return repo
}

func (r Repo) InsertEarthquake(earthquake models.Earthquake) error {
	stmt, err := r.Db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(earthquake.Lat, earthquake.Lon, earthquake.Mag)
	if err != nil{
		return err
	}
	return nil
}
