package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int
	DatabaseURL string
	JWTSecret   string
	Env         string
}

var Cfg *Config

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.AutomaticEnv()

	Cfg = &Config{
		Port:        viper.GetInt("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
		Env:         viper.GetString("ENV"),
	}

	fmt.Println("Config loaded successfully")
}
