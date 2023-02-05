package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *ConfigStruct
)

type ConfigStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

func ReadConfig(path string) (*ConfigStruct, error) {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix

	return config, nil

}
