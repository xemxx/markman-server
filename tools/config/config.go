package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		JwtSecret        string `mapstructure:"jwt_secret"`
		RunMode          string `mapstructure:"run_mode"`
		Version          string `mapstructure:"version"`
		MinClientVersion string `mapstructure:"min_client_version"`
	} `mapstructure:"app"`
	Server struct {
		HttpPort     string        `mapstructure:"http_port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	} `mapstructure:"server"`
	Database struct {
		Type     string `mapstructure:"type"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Database string `mapstructure:"database"`
	} `mapstructure:"database"`
	Runtime struct {
		LogUrl string `mapstructure:"log_url"`
	} `mapstructure:"runtime"`
}

var Cfg *Config

func Init(configFile string) {
	c := viper.New()
	c.SetConfigType("yaml")
	c.SetConfigFile(configFile)
	err := c.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error when reading config file: %s\n", err))
	}
	if err = c.Unmarshal(&Cfg); err != nil {
		panic(fmt.Sprintf("Fatal error when unmarshal config file: %s\n", err))
	}
}
