package tvdb

type loginAPIResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
}

type nextAiredData struct {
	ID                   int64    `json:"id"`
	Name                 string   `json:"name"`
	Slug                 string   `json:"slug"`
	Image                string   `json:"image"`
	NameTranslations     []string `json:"nameTranslations"`
	OverviewTranslations []string `json:"overviewTranslations"`
	Aliases              []struct {
		Language string `json:"language"`
		Name     string `json:"name"`
	} `json:"aliases"`
	FirstAired string `json:"firstAired"`
	LastAired  string `json:"lastAired"`
	NextAired  string `json:"nextAired"`
	Score      int64  `json:"score"`
	Status     struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		RecordType  string `json:"recordType"`
		KeepUpdated bool   `json:"keepUpdated"`
	} `json:"status"`
	OriginalCountry   string      `json:"originalCountry"`
	OriginalLanguage  string      `json:"originalLanguage"`
	DefaultSeasonType int64       `json:"defaultSeasonType"`
	IsOrderRandomized bool        `json:"isOrderRandomized"`
	LastUpdated       string      `json:"lastUpdated"`
	AverageRuntime    int64       `json:"averageRuntime"`
	Episodes          interface{} `json:"episodes"`
	Overview          string      `json:"overview"`
	Year              string      `json:"year"`
}

type nextAiredResponse struct {
	Status string        `json:"status"`
	Data   nextAiredData `json:"data"`
}
