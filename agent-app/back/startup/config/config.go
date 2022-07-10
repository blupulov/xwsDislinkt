package config

import "os"

type Config struct {
	RestPort       string
	AgentAppDBHost string
	AgentAppDBPort string
	GrpcPort       string
}

func NewConfig() *Config {
	return &Config{
		RestPort:       os.Getenv("AGENT_APP_PORT"),
		GrpcPort:       "9053",
		AgentAppDBHost: os.Getenv("MONGO_DB_HOST"),
		AgentAppDBPort: os.Getenv("MONGO_DB_PORT"),
	}
}
