package models

type RepoConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  string
	Broker   string
	Topic    string
}

type Earthquake struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Mag float64 `json:"mag"`
}
