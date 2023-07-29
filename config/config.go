package config

const (
	FlightStopMaxNumber = 5

	DbHost     = "localhost"
	DbPort     = 8080
	DbName     = "aviato"
	DbUser     = "postgres"
	DbPassword = "778977"
)

type StorageConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig() StorageConfig {

	return StorageConfig{
		Host:     DbHost,
		Port:     DbPort,
		Database: DbName,
		Username: DbUser,
		Password: DbPassword,
	}
}
