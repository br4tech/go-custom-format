package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
	}

	App struct {
		Port int
	}
)

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v", err))
	}

	return Config{
		App: App{
			Port: viper.GetInt("app.server.port"),
		},
	}
}
