package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"script/config"
	"script/tracked"
	"script/tvdb"
	"time"

	"golang.org/x/exp/slog"
)

var opts = slog.HandlerOptions{
	Level: slog.LevelDebug, //TODO: Get from conf object above. More details in config/config.go
}
var textHandler = opts.NewTextHandler(os.Stdout)
var logger = slog.New(textHandler)

func main() {
	logger.Info("Starting script")
	conf, err := config.ReadConfig()

	if err != nil {
		logger.Error("Error after reading config", err)
		fmt.Println(err.Error())
		return
	}

	logger.Debug("Initiating tvdb client and logging in")
	tv := tvdb.Client{ApiKey: conf.TvdbApiKey, Pin: conf.TvdbPin}
	err = tv.Login()
	if err != nil {
		logger.Error("Error in tvdb login", err)
		panic(err)
	}

	logger.Debug("Getting tracked tv shows from database")
	tracker := tracked.TrackedService{}
	shows := tracker.GetTrackedTvShows()

	logger.Debug("Getting days left using tvdb api")
	for _, show := range shows {
		logger.Debug("Getting days left for ", show.Title)
		resp := getDaysLeftResponseForTvShow(&tv, show.TvdbId)
		logger.Debug("initiating Discord webhook call")
		discordWebhookPost(conf.DiscordWebhookUrl, resp)

	}
	logger.Info("Complete")
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
	if err != nil {
		logger.Error("Error in discord webhook post", err)
		panic(err)
	}

	body, err := io.ReadAll(response.Body)
	if response.StatusCode != 200 && response.StatusCode != 204 {
		logger.Error(fmt.Sprintf("Error in discord webhook post, got status: %d and response body: %s", response.StatusCode, string(body)), nil)
		panic(err)
	}

}
