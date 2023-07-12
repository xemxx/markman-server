package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// todo 转换为结构体加yaml
type config struct {
	*viper.Viper
}

var Cfg *config

func Init(configFile string) {
	Cfg = &config{
		viper.New(),
	}
	Cfg.SetConfigFile(configFile)
	err := Cfg.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Fatal error when reading config file: %s\n", err))
	}
}
