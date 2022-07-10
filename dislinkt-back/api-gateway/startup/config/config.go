package config

import "os"

type Config struct {
	Port                 string
	PostServiceHost      string
	PostServicePort      string
	UserServiceHost      string
	UserServicePort      string
	FollowingServiceHost string
	FollowingServicePort string
	CompanyServiceHost   string
	CompanyServicePort   string
}

//dodati u .env
func NewConfig() *Config {
	return &Config{
		Port:                 os.Getenv("GATEWAY_PORT"),
		PostServiceHost:      os.Getenv("POST_SERVICE_HOST"),
		PostServicePort:      os.Getenv("POST_SERVICE_PORT"),
		UserServiceHost:      os.Getenv("USER_SERVICE_HOST"),
		UserServicePort:      os.Getenv("USER_SERVICE_PORT"),
		FollowingServiceHost: os.Getenv("FOLLOWING_SERVICE_HOST"),
		FollowingServicePort: os.Getenv("FOLLOWING_SERVICE_PORT"),
		CompanyServiceHost:   "localhost",
		CompanyServicePort:   "9053",
	}
}
