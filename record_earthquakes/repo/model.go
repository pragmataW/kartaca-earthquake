package repo

import (
	"database/sql"
	"sync"
)

type Repo struct {
	Db     *sql.DB
}

var (
	repo *Repo
	once sync.Once
)

const (
	insertQuery = "INSERT INTO earthquakes (lat, lon, magnitude) VALUES ($1, $2, $3)"
)
