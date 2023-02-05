package tracked

import (
	"database/sql"
	"log"
	"script/config"

	_ "github.com/go-sql-driver/mysql"
)

type TrackedService struct{}

func (t *TrackedService) GetTrackedTvShows() []TvShow {
	conf, _ := config.ReadConfig()
	db, err := sql.Open("mysql", conf.MysqlUri)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * from tracked_shows")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	var shows []TvShow
	for res.Next() {
		var show TvShow
		err := res.Scan(&show.TvdbId, &show.Title)

		if err != nil {
			log.Fatal(err)
		}
		shows = append(shows, show)
	}
	// shows := []TvShow{
	// 	{TvdbId: "359274", Title: "Vinland Saga"},
	// 	// {TvdbId: "67890"},
	// }
	return shows
}
