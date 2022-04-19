package config

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8050",
		PostDBHost: "localhost",
		PostDBPort: "27017",
	}
}
