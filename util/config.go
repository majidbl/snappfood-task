package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// Viper reads the values from a config file or environment variable.
type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	QueueType         string `mapstructure:"QUEUE_TYPE"`      // default memory: memory, redis
	OrderQueueKey     string `mapstructure:"ORDER_QUEUE_KEY"` // default memory: memory, redis
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
