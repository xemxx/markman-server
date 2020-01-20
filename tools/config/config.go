package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	*viper.Viper
}

var Cfg *config

func init(){
	Cfg = &config{
		viper.New(),
	}
	Cfg.SetConfigFile("app.json")
	err := Cfg.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Fatal error when reading config file: %s\n", err))
	}
}
