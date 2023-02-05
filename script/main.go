package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"script/config"
	"script/tvdb"
	"time"
)

func main() {
	conf, err := config.ReadConfig("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(conf.DiscordWebhookUrl)

	tv := tvdb.Client{ApiKey: conf.TvdbApiKey, Pin: conf.TvdbPin}
	err = tv.Login()
	if err != nil {
		panic(err)
	}
	resp := getDaysLeftResponseForTvShow(&tv, "359274")
	discordWebhookPost(conf.DiscordWebhookUrl, resp)
}

func getDaysLeftResponseForTvShow(tv *tvdb.Client, seriesId string) WebhookBody {
	data, err := tv.GetSeriesNextAiredResponse("359274")
	if err != nil {
		panic(err)
	}
	timeFormat := "2006-01-02"
	nextAired, _ := time.Parse(data.Data.NextAired, timeFormat)
	hoursLeft := nextAired.Sub(time.Now()).Hours()
    fmt.Printf("Hours left: %d\n", &hoursLeft)
	daysLeft := int(hoursLeft / 24)

	resp := WebhookBody{
		Content: fmt.Sprintf("% days left for %s", daysLeft, data.Data.Name),
		Embeds: []Embed{
			{
				Image: Image{
					URL: data.Data.Image,
				},
			},
		},
	}
	return resp

}
func discordWebhookPost(discordWebhookUrl string, params WebhookBody) {
	jsonMarshal, _ := json.Marshal(params)

	response, err := http.Post(discordWebhookUrl, "application/json", bytes.NewBuffer(jsonMarshal))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
