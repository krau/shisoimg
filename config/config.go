package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"host" toml:"host"`
}

var cfg *Config

func LoadConfig() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		if _, err := os.Create("config.toml"); err != nil {
			return nil, err
		}

	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SHISOIMG")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("host", ":34180")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func Host() string {
	cfg, err := LoadConfig()
	if err != nil {
		return ""
	}

	return cfg.Host
}
