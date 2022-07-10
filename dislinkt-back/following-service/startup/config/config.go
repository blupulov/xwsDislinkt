package config

import "os"

type Config struct {
	RestPort         string
	GrpcPort         string
	FollowDBHost     string
	FollowDBPort     string
	FollowDBUsername string
	FollowDBPassword string
}

func NewConfig() *Config {
	return &Config{
		RestPort:         "8052",
		GrpcPort:         os.Getenv("FOLLOWING_SERVICE_PORT"),
		FollowDBHost:     os.Getenv("NEO4J_DB_HOST"),
		FollowDBPort:     os.Getenv("NEO4J_DB_PORT"),
		FollowDBUsername: "neo4j",
		FollowDBPassword: "password",
	}
}
