package config

type Config struct {
	RestPort      string
	CompanyDBHost string
	CompanyDBPort string
	GrpcPort      string
}

func NewConfig() *Config {
	return &Config{
		RestPort:      "8053",
		GrpcPort:      "9053",
		CompanyDBHost: "localhost",
		CompanyDBPort: "27017",
	}
}
