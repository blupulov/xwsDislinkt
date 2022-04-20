package config

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8051",
		UserDBHost: "localhost",
		UserDBPort: "27017",
	}
}
