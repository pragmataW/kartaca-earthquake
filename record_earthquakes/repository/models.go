package repository

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	repo *Repo
	once sync.Once
)

const (
	insertQuery = "INSERT INTO earthquakes (lat, lon, magnitude) VALUES ($1, $2, $3)"
)

type Repo struct {
	Db     *sql.DB
	Broker string
	Topic  string
}