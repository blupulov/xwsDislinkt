package config

type Config struct {
	RestPort   string
	PostDBHost string
	PostDBPort string
	GrpcPort   string
}

func NewConfig() *Config {
	return &Config{
		RestPort:   "8050",
		GrpcPort:   "9050",
		PostDBHost: "localhost",
		PostDBPort: "27017",
	}
}
