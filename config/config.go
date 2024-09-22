package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServiceName string `MAPSTRUCTURE:"SERVICE_NAME"`
	Environment string `MAPSTRUCTURE:"ENVIRONMENT"`
	Version     string `MAPSTRUCTURE:"VERSION"`
	Port        string `MAPSTRUCTURE:"PORT"`

	Database DatabaseConfig `MAPSTRUCTURE:"DATABASE"`
	Redis    RedisConfig    `MAPSTRUCTURE:"REDIS"`
}

type AppSecret struct {
	JWTSecretKey string         `MAPSTRUCTURE:"JWT_SECRET_KEY"`
	Database     DatabaseSecret `MAPSTRUCTURE:"DATABASE"`
	Redis        RedisSecret    `MAPSTRUCTURE:"REDIS"`
}

type DatabaseConfig struct {
	Host     string `MAPSTRUCTURE:"HOST"`
	Port     string `MAPSTRUCTURE:"PORT"`
	Database string `MAPSTRUCTURE:"DATABASE"`
}

type DatabaseSecret struct {
	Username string `MAPSTRUCTURE:"USERNAME"`
	Password string `MAPSTRUCTURE:"PASSWORD"`
}

type RedisConfig struct {
	Addr string `MAPSTRUCTURE:"ADDR"`
	DB   int    `MAPSTRUCTURE:"DB"`
}

type RedisSecret struct {
	Password string `MAPSTRUCTURE:"PASSWORD"`
}

func New(configPath, secretPath string) (*AppConfig, *AppSecret, error) {
	// Load Config
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading config file: %w", err)
	}
	var config *AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	viper.Reset()

	// Load Secret
	viper.SetConfigFile(secretPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading secret file: %w", err)
	}
	var secret *AppSecret
	if err := viper.Unmarshal(&secret); err != nil {
		return nil, nil, fmt.Errorf("error unmarshalling secret: %w", err)
	}

	return config, secret, nil
}
