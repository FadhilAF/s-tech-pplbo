package config

import (
  "log"

	"github.com/spf13/viper"
)

type Config struct {
  AppPort string `mapstructure:"APP_PORT"`
  PostgresUrl string `mapstructure:"POSTGRES_CONNECTION_URL"` 
}

func LoadConfig(envPath string) Config {
  var config Config

  viper.SetConfigFile(envPath)

  err := viper.ReadInConfig()

  if err != nil {
    log.Fatalf("Error occured while .env reading: %v", err)
  }

  viper.SetDefault("PORT", "8000")
  
  viper.Unmarshal(&config)

  return config
}
