package overseerr_structs

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductionCompany struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
}

type ProductionCountry struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

type SpokenLanguage struct {
	EnglishName string `json:"englishName"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
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

type MediaRequest struct {
	ID         int         `json:"id"`
	Status     int         `json:"status"`
	Media      MediaInfo   `json:"media"`
	CreatedAt  string      `json:"createdAt"`
	UpdatedAt  string      `json:"updatedAt"`
	RequestBy  User        `json:"requestedBy"`
	ModifiedBy User        `json:"modifiedBy"`
	Is4k       bool        `json:"is4k"`
	ServerID   interface{} `json:"serverId"`
	ProfileID  interface{} `json:"profileId"`
	RootFolder interface{} `json:"rootFolder"`
	RatingKey  string      `json:"ratingKey"`
	Seasons    []struct {
		ID           int    `json:"id"`
		SeasonNumber int    `json:"seasonNumber"`
		Status       int    `json:"status"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
	} `json:"seasons"`
}

type MediaInfo struct {
	ID        int            `json:"id"`
	MediaType string         `json:"mediaType"`
	TmdbID    int            `json:"tmdbId"`
	TvdbID    int            `json:"tvdbId"`
	Status    int            `json:"status"`
	Requests  []MediaRequest `json:"requests"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	RatingKey string         `json:"ratingKey"`
	Seasons   []struct {
		ID           int    `json:"id"`
		SeasonNumber int    `json:"seasonNumber"`
		Status       int    `json:"status"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
	} `json:"seasons"`
}

type ExternalIds struct {
	FacebookID  interface{} `json:"facebookId"`
	FreebaseID  interface{} `json:"freebaseId"`
	FreebaseMid interface{} `json:"freebaseMid"`
	ImdbID      string      `json:"imdbId"`
	InstagramID interface{} `json:"instagramId"`
	TvdbID      int         `json:"tvdbId"`
	TvrageID    interface{} `json:"tvrageId"`
	TwitterID   interface{} `json:"twitterId"`
}

type WatchProviderDetails struct {
	DisplayPriority int    `json:"displayPriority"`
	LogoPath        string `json:"logoPath"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
}

// Missing FlatRate
type WatchProvider struct {
	Iso3166_1 string                 `json:"iso_3166_1"`
	Link      string                 `json:"link"`
	Buy       []WatchProviderDetails `json:"buy"`
	FlatRate  []struct{}             `json:"flatrate"`
}
