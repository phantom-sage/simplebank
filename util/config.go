package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config hold the configuration for the application.
type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	AllowedOrigins       []string      `mapstructure:"ALLOWED_ORIGINS"`
	DBConnUrl            string        `mapstructure:"DB_CONN_URL"`
	MigrationUrl         string        `mapstructure:"MIGRATION_URL"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress    string        `mapstructure:"GRPc_SERVER_ADDRESS"`
	TokenSymmetricKEy    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

// LoadConfig load config from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
