package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"script/config"
	"strings"
)

func main() {
	conf, err := config.ReadConfig("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(conf.DiscordWebhookUrl)
	requestBody := strings.NewReader(`
		{
			"content":"Hello There!"
		}
	`)

	response, err := http.Post(conf.DiscordWebhookUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
