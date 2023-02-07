package config

import (
	"fmt"

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
	MysqlUri          string `json: "MysqlUri"`
	// LogLevel          slog.Level `json : "LogLevel"` //TODO: I think Unmarshal should have been implemented so it should be parsing string like DEBUG/INFO from env and parsing the right log level
}

func ReadConfig() (*ConfigStruct, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	//TODO: There has to be a better way to override config.json via environment than below
	viper.BindEnv("DiscordWebhookUrl", "DiscordWebhookUrl")
	viper.BindEnv("TvdbApiKey", "TvdbApiKey")
	viper.BindEnv("LogLevel", "LogLevel")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("before unmarshal")
	err = viper.Unmarshal(&config)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return config, nil

}
