package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"script/config"
	"script/tracked"
	"script/tvdb"
	"time"
)

func main() {
	conf, err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tracker := tracked.TrackedService{}
	tv := tvdb.Client{ApiKey: conf.TvdbApiKey, Pin: conf.TvdbPin}
	err = tv.Login()
	if err != nil {
		panic(err)
	}
	shows := tracker.GetTrackedTvShows()
	for _, show := range shows {
		resp := getDaysLeftResponseForTvShow(&tv, show.TvdbId)
		discordWebhookPost(conf.DiscordWebhookUrl, resp)

	}
}

func getDaysLeftResponseForTvShow(tv *tvdb.Client, seriesId string) WebhookBody {
	data, err := tv.GetSeriesNextAiredResponse(seriesId)
	if err != nil {
		panic(err)
	}
	timeFormat := "2006-01-02"
	nextAired, _ := time.Parse(timeFormat, data.Data.NextAired)
	daysLeft := timeDiff(time.Now(), nextAired)

	resp := WebhookBody{
		Content: fmt.Sprintf("%d days left for %s", daysLeft, data.Data.Name),
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

func timeDiff(source, dest time.Time) int {
	hoursLeft := dest.Sub(source).Hours()
	return int(hoursLeft / 24)
}
func discordWebhookPost(discordWebhookUrl string, params WebhookBody) {
	jsonMarshal, _ := json.Marshal(params)

	response, err := http.Post(discordWebhookUrl, "application/json", bytes.NewBuffer(jsonMarshal))

	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

}
