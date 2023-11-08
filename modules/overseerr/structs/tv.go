package overseerr_structs

type Episode struct {
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
	ID           int       `json:"id"`
	AirDate      string    `json:"airDate"`
	EpisodeCount int       `json:"episodeCount"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	PosterPath   string    `json:"posterPath"`
	SeasonNumber int       `json:"seasonNumber"`
	Episodes     []Episode `json:"episodes"`
}

type TvResult struct {
	ID               int       `json:"id"`
	MediaType        string    `json:"mediaType"`
	Popularity       float64   `json:"popularity"`
	PosterPath       string    `json:"posterPath"`
	BackdropPath     string    `json:"backdropPath"`
	VoteCount        int       `json:"voteCount"`
	VoteAverage      float64   `json:"voteAverage"`
	GenreIds         []int     `json:"genreIds"`
	Overview         string    `json:"overview"`
	OriginalLanguage string    `json:"originalLanguage"`
	Name             string    `json:"name"`
	OriginalName     string    `json:"originalName"`
	OriginCountry    []string  `json:"originCountry"`
	FirstAirDate     string    `json:"firstAirDate"`
	MediaInfo        MediaInfo `json:"mediaInfo"`
}

type ContentRating struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Rating    string `json:"rating"`
}

type Creator struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profilePath"`
}

type Network struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
}

type Credits struct {
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TvDetails struct {
	ID                  int                 `json:"id"`
	BackdropPath        string              `json:"backdropPath"`
	PosterPath          string              `json:"posterPath"`
	ContentRatings      ContentRating       `json:"contentRatings"`
	CreatedBy           []Creator           `json:"createdBy"`
	EpisodeRunTime      []int               `json:"episodeRunTime"`
	FirstAirDate        string              `json:"firstAirDate"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	InProduction        bool                `json:"inProduction"`
	Languages           []string            `json:"languages"`
	LastAirDate         string              `json:"lastAirDate"`
	LastEpisodeToAir    Episode             `json:"lastEpisodeToAir"`
	Name                string              `json:"name"`
	NextEpisodeToAir    Episode             `json:"nextEpisodeToAir"`
	Networks            []Network           `json:"networks"`
	NumberOfEpisodes    int                 `json:"numberOfEpisodes"`
	NumberOfSeasons     int                 `json:"numberOfSeason"`
	OriginCountry       []string            `json:"originCountry"`
	OriginalLanguage    string              `json:"originalLanguage"`
	OriginalName        string              `json:"originalName"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	ProductionCompanies []ProductionCompany `json:"productionCompanies"`
	ProductionCountries []ProductionCountry `json:"productionCountries"`
	SpokenLanguages     []SpokenLanguage    `json:"spokenLanguages"`
	Seasons             []Season            `json:"seasons"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"voteAverage"`
	VoteCount           int                 `json:"voteCount"`
	Credits         struct {
		Cast []Cast `json:"cast"`
		Crew []Crew `json:"crew"`
	} `json:"credits"`
	ExternalIds         ExternalIds         `json:"externalIds"`
	Keywords            []Keyword           `json:"keywords"`
	MediaInfo           MediaInfo           `json:"mediaInfo"`
	WatchProviders      []WatchProvider     `json:"watchProviders"`
}
