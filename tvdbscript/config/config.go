package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	Token     string
	BotPrefix string

	config *ConfigStruct
)

type ConfigStruct struct {
	DiscordWebhookUrl string `json : "DiscordWebhookUrl"`
	TvdbApiKey        string `json : "TvdbApiKey"`
	TvdbPin           string `json : "TvdbPin"`
}

func ReadConfig(path string) (*ConfigStruct, error) {
    fmt.Println(os.Getenv("TvdbApiKey"))
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.BindEnv("TvdbApiKey", "TvdbApiKey")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}
	err = viper.Unmarshal(&config)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return config, nil

}
