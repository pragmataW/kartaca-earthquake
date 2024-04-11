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
	Lat float64
	Lon float64
	Mag float64
}
