package config

import "os"

type Config struct {
	RestPort   string
	UserDBHost string
	UserDBPort string
	GrpcPort   string
}

func NewConfig() *Config {
	return &Config{
		RestPort:   "8051",
		GrpcPort:   os.Getenv("USER_SERVICE_PORT"),
		UserDBHost: os.Getenv("MONGO_DB_HOST"),
		UserDBPort: os.Getenv("MONGO_DB_PORT"),
	}
}
