package configs

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"time"
)

func init() {
	value, err := getEnv("FILE_CONFIG_NAME")
	if err != nil {
		event := log.Panic()
		event.Msgf("error reading env key: FILE_CONFIG_NAME")
	}
	log.Info().Msgf("config file is set to: %s", value)
	viper.SetConfigName(value)
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	err = viper.ReadInConfig()
	if err != nil {
		event := log.Panic()
		event.Msgf("Fatal error config file: %w", err)
	}
}

func GracefulTimeout() time.Duration {
	return viper.GetDuration("service.gracefulShutdownTimeout")
}

func DBName() string {
	return viper.GetString("mongo.dbName")
}

func ServiceAddress() string {
	return viper.GetString("service.address")
}

func MongoURI() string {
	return viper.GetString("mongo.uri")
}

func getEnv(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", fmt.Errorf("key not found: %s", key)
}
