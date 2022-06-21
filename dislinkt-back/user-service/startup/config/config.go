package config

type Config struct {
	RestPort   string
	UserDBHost string
	UserDBPort string
	GrpcPort   string
}

func NewConfig() *Config {
	return &Config{
		RestPort:   "8051",
		GrpcPort:   "9051",
		UserDBHost: "localhost",
		UserDBPort: "27017",
	}
}
