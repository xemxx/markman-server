package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	//Cfg .
	Cfg *viper.Viper
)

func init() {
	Cfg = viper.New()
	Cfg.SetConfigFile("config/app.json")
	err := Cfg.ReadInConfig()

	if err != nil {
		fmt.Printf("Fatal error when reading config file: %s\n", err)
	}
}

// Retu .
func Retu(str string) string {
	return Cfg.GetString(str)
}
