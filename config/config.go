package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	InfuraURL   string
	BlockNumber uint64

	Timeout int
	Host    string
	Port    string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RateLimitCapacity int
	OtpExpireInMins   int

	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDB       int

	SmtpHost     string
	SmtpPort     int
	SmtpUsername string
	SmtpPassword string
	SmtpFrom     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	viper.AutomaticEnv()

	return &Config{
		InfuraURL:   viper.GetString("INFURA_URL"),
		BlockNumber: viper.GetUint64("BLOCK_NUMBER"),

		Timeout: viper.GetInt("TIMEOUT"),
		Host:    viper.GetString("HOST"),
		Port:    viper.GetString("PORT"),

		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),

		RateLimitCapacity: viper.GetInt("RATE_LIMIT_CAPACITY"),
		OtpExpireInMins:   viper.GetInt("OTP_EXPIRE_IN_MINS"),

		RedisHost:     viper.GetString("REDIS_HOST"),
		RedisPort:     viper.GetInt("REDIS_PORT"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisDB:       viper.GetInt("REDIS_DB"),

		SmtpHost:     viper.GetString("SMTP_HOST"),
		SmtpPort:     viper.GetInt("SMTP_PORT"),
		SmtpUsername: viper.GetString("SMTP_USERNAME"),
		SmtpPassword: viper.GetString("SMTP_PASSWORD"),
		SmtpFrom:     viper.GetString("SMTP_FROM"),
	}
}
