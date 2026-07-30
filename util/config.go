package util

import "github.com/spf13/viper"

// Config hold the configuration for the application.
type Config struct {
	DBConnUrl     string `mapstructure:"DB_CONN_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig load config from file or environment variables.
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
