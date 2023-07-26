package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUri                  string        `mapstruct:"MONGO_URI"`
	RedisUri               string        `mapstruct:"REDIS_URL"`
	Port                   string        `mapstruct:"PORT"`
	AccessTokenPrivateKey  string        `mapstruct:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstruct:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstruct:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstruct:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstruct:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstruct:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstruct:"ACCESS_TOKEN_MAX_AGE"`
	RefreshTokenMaxAge     int           `mapstruct:"REFRESH_TOKEN_MAX_AGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
