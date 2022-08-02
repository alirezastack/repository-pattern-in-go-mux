package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"time"
)

func init() {
	log.Info().Msgf("config file is set to: %s", getEnv("FILE_CONFIG_NAME", ""))
	viper.SetConfigName(getEnv("FILE_CONFIG_NAME", ""))
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
