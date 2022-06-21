package config

type Config struct {
	RestPort       string
	AgentAppDBHost string
	AgentAppDBPort string
	GrpcPort       string
}

func NewConfig() *Config {
	return &Config{
		RestPort:       "8053",
		GrpcPort:       "9053",
		AgentAppDBHost: "localhost",
		AgentAppDBPort: "27017",
	}
}
