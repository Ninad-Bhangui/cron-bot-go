package tracked

type TrackedService struct{}

func (t *TrackedService) GetTrackedTvShows() []TvShow {
	shows := []TvShow{
        {TvdbId: "359274", Title: "Vinland Saga"},
		// {TvdbId: "67890"},
	}
	return shows
}
