package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Token    string         `mapstructure:"token"    yaml:"token"`
	Postgres PostgresConfig `mapstructure:"postgres" yaml:"postgres"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile("./config/config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}
