package config

import (
	"log"

	"github.com/spf13/viper"
)

// Adjust cfg type as needed
type Cfg struct {
	Mp mailpit
}

type mailpit struct {
	ApiUrl       string
	WebsocketUrl string
}

func InitConfig() *Cfg {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	return &Cfg{
		Mp: mailpit{
			ApiUrl:       viper.GetString("mailpit.api_url"),
			WebsocketUrl: viper.GetString("mailpit.websocket_url"),
		},
	}
}
