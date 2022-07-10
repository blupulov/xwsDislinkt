package config

import "os"

type Config struct {
	RestPort   string
	PostDBHost string
	PostDBPort string
	GrpcPort   string
}

func NewConfig() *Config {
	return &Config{
		RestPort:   "8050",
		GrpcPort:   os.Getenv("POST_SERVICE_PORT"),
		PostDBHost: os.Getenv("MONGO_DB_HOST"),
		PostDBPort: os.Getenv("MONGO_DB_PORT"),
	}
}
