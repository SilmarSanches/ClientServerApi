package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	URLDolar string `mapstructure:"URL_EXCHANGE"`
	Database string `mapstructure:"DATABASE"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()

	_ = viper.BindEnv("URL_EXCHANGE")
	_ = viper.BindEnv("DATABASE")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	_ = os.Setenv("URL_EXCHANGE", cfg.URLDolar)
	_ = os.Setenv("DATABASE", cfg.Database)

	return &cfg, nil
}
