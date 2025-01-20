package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	UrlApi string `mapstructure:"URL_API"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()

	_ = viper.BindEnv("URL_API")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	_ = os.Setenv("URL_API", cfg.UrlApi)

	return &cfg, nil
}
