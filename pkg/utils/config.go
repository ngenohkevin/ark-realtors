package utils

import (
	"errors"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DbDriver             string        `mapstructure:"DB_DRIVER"`
	DbUrl                string        `mapstructure:"DB_URL"`
	ServerAddr           string        `mapstructure:"SERVER_ADDR"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig loads the configuration from the environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return Config{}, errors.New("config file \".env\" not found")
		}
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	return
}
