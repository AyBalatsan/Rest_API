package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}
}

// функция для добавление конфигов
func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to load config file")
	}
}

func ConfigurationInstaller() {
	initConfig()
	loadEnv()
}
