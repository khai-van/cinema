package config

import (
	"strings"

	"github.com/spf13/viper"
)

var config Config

type ServiceConfig struct {
	Port int `mapstructure:"port"`

	MinDistance int `mapstructure:"min_distance"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	MaxIdleConns     int    `mapstructure:"max_idle_connections"`
	MaxOpenConns     int    `mapstructure:"max_open_connections"`
}

type Config struct {
	Service  ServiceConfig  `mapstructure:"service"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

func LoadConfig(filePath, fileName string) (*Config, error) {
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filePath)

	// Automatically bind environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func Get() *Config {
	return &config
}
