package tv

type ContentRating struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Rating    string `json:"rating"`
}

type CreatedBy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profilePath"`
}

type LastEpisode struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	AirDate        string  `json:"airDate"`
	EpisodeNumber  int     `json:"episodeNumber"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"productionCode"`
	SeasonNumber   int     `json:"seasonNumber"`
	ShowID         int     `json:"showId"`
	StillPath      string  `json:"stillPath"`
	VoteAverage    float64 `json:"voteAverage"`
	VoteCount      int     `json:"voteCount"`
}

type Season struct {
	ID           int           `json:"id"`
	AirDate      string        `json:"airDate"`
	EpisodeCount int           `json:"episodeCount"`
	Name         string        `json:"name"`
	Overview     string        `json:"overview"`
	PosterPath   string        `json:"posterPath"`
	SeasonNumber int           `json:"seasonNumber"`
	Episodes     []LastEpisode `json:"episodes"`
}

type Cast struct {
	ID          int    `json:"id"`
	CastID      int    `json:"castId"`
	Character   string `json:"character"`
	CreditID    string `json:"creditId"`
	Gender      int    `json:"gender"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	ProfilePath string `json:"profilePath"`
}

type Crew struct {
	ID          int    `json:"id"`
	CreditID    string `json:"creditId"`
	Gender      int    `json:"gender"`
	Name        string `json:"name"`
	Job         string `json:"job"`
	Department  string `json:"department"`
	ProfilePath string `json:"profilePath"`
}

type ExternalIds struct {
	FacebookID  string `json:"facebookId"`
	FreebaseID  string `json:"freebaseId"`
	FreebaseMid string `json:"freebaseMid"`
	ImdbID      string `json:"imdbId"`
	InstagramID string `json:"instagramId"`
	TvdbID      int    `json:"tvdbId"`
	TvrageID    int    `json:"tvrageId"`
	TwitterID   string `json:"twitterId"`
}

type MediaRequester struct {
	ID     int `json:"id"`
	Status int `json:"status"`
	Media  struct {
		DownloadStatus        []interface{} `json:"downloadStatus"`
		DownloadStatus4k      []interface{} `json:"downloadStatus4k"`
		ID                    int           `json:"id"`
		MediaType             string        `json:"mediaType"`
		TmdbID                int           `json:"tmdbId"`
		TvdbID                int           `json:"tvdbId"`
		ImdbID                string        `json:"imdbId"`
		Status                int           `json:"status"`
		Status4k              int           `json:"status4k"`
		CreatedAt             string        `json:"createdAt"`
		UpdatedAt             string        `json:"updatedAt"`
		LastSeasonChange      string        `json:"lastSeasonChange"`
		MediaAddedAt          string        `json:"mediaAddedAt"`
		ServiceID             int           `json:"serviceId"`
		ServiceID4k           int           `json:"serviceId4k"`
		ExternalServiceID     string        `json:"externalServiceId"`
		ExternalServiceID4k   string        `json:"externalServiceId4k"`
		ExternalServiceSlug   string        `json:"externalServiceSlug"`
		ExternalServiceSlug4k string        `json:"externalServiceSlug4k"`
		RatingKey             string        `json:"ratingKey"`
		RatingKey4k           string        `json:"ratingKey4k"`
	} `json:"media"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	RequestedBy struct {
		ID           int    `json:"id"`
		Email        string `json:"email"`
		Username     string `json:"username"`
		PlexToken    string `json:"plexToken"`
		PlexUsername string `json:"plexUsername"`
		UserType     int    `json:"userType"`
		Permissions  int    `json:"permissions"`
		Avatar       string `json:"avatar"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		RequestCount int    `json:"requestCount"`
	} `json:"requestedBy"`
	ModifiedBy struct {
		ID           int    `json:"id"`
		Email        string `json:"email"`
		Username     string `json:"username"`
		PlexToken    string `json:"plexToken"`
		PlexUsername string `json:"plexUsername"`
		UserType     int    `json:"userType"`
		Permissions  int    `json:"permissions"`
		Avatar       string `json:"avatar"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		RequestCount int    `json:"requestCount"`
	} `json:"modifiedBy"`
	Is4k       bool   `json:"is4k"`
	ServerID   int    `json:"serverId"`
	ProfileID  int    `json:"profileId"`
	RootFolder string `json:"rootFolder"`
}

type Link struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Link      string `json:"link"`
	Buy       []struct {
		DisplayPriority int    `json:"displayPriority"`
		LogoPath        string `json:"logoPath"`
		ID              int    `json:"id"`
		Name            string `json:"name"`
	} `json:"buy"`
	FlatRate []struct {
		DisplayPriority int    `json:"displayPriority"`
		LogoPath        string `json:"logoPath"`
		ID              int    `json:"id"`
		Name            string `json:"name"`
	} `json:"flatrate"`
}

type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tv struct {
	ID             int           `json:"id"`
	BackdropPath   string        `json:"backdropPath"`
	PosterPath     string        `json:"posterPath"`
	ContentRatings ContentRating `json:"contentRatings"`
	CreatedBy      []CreatedBy   `json:"createdBy"`
	EpisodeRunTime []int         `json:"episodeRunTime"`
	FirstAirDate   string        `json:"firstAirDate"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string      `json:"homepage"`
	InProduction     bool        `json:"inProduction"`
	Languages        []string    `json:"languages"`
	LastAirDate      string      `json:"lastAirDate"`
	LastEpisodeToAir LastEpisode `json:"lastEpisodeToAir"`
	Name             string      `json:"name"`
	NextEpisodeToAir LastEpisode `json:"nextEpisodeToAir"`
	Networks         []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logoPath"`
		OriginCountry string `json:"originCountry"`
		Name          string `json:"name"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"numberOfEpisodes"`
	NumberOfSeasons     int      `json:"numberOfSeason"`
	OriginCountry       []string `json:"originCountry"`
	OriginalLanguage    string   `json:"originalLanguage"`
	OriginalName        string   `json:"originalName"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logoPath"`
		OriginCountry string `json:"originCountry"`
		Name          string `json:"name"`
	} `json:"productionCompanies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
	} `json:"productionCountries"`
	SpokenLanguages []struct {
		EnglishName string `json:"englishName"`
		Iso639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spokenLanguages"`
	Seasons     []Season `json:"seasons"`
	Status      string   `json:"status"`
	Tagline     string   `json:"tagline"`
	Type        string   `json:"type"`
	VoteAverage float64  `json:"voteAverage"`
	VoteCount   int      `json:"voteCount"`
	Credits     struct {
		Cast []Cast `json:"cast"`
		Crew []Crew `json:"crew"`
	} `json:"credits"`
	ExternalIds ExternalIds `json:"externalIds"`
	Keywords    []Keyword   `json:"keywords"`
	MediaInfo   struct {
		ID        int              `json:"id"`
		TmdbID    int              `json:"tmdbId"`
		TvdbID    int              `json:"tvdbId"`
		Status    int              `json:"status"`
		Requests  []MediaRequester `json:"requests"`
		CreatedAt string           `json:"createdAt"`
		UpdatedAt string           `json:"updatedAt"`
	} `json:"mediaInfo"`
	WatchProviders []Link `json:"watchProviders"`
}
