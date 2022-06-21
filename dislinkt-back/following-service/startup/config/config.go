package config

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
		GrpcPort:         "9052",
		FollowDBHost:     "localhost",
		FollowDBPort:     "7687",
		FollowDBUsername: "neo4j",
		FollowDBPassword: "password",
	}
}
