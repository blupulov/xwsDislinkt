package config

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
		Port:                 "8001",
		PostServiceHost:      "localhost",
		PostServicePort:      "9050",
		UserServiceHost:      "localhost",
		UserServicePort:      "9051",
		FollowingServiceHost: "localhost",
		FollowingServicePort: "9052",
		CompanyServiceHost:   "localhost",
		CompanyServicePort:   "9053",
	}
}
