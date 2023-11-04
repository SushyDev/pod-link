package overseerr

import "time"

type User struct {
	ID           int         `json:"id"`
	Email        string      `json:"email"`
	Username     interface{} `json:"username"`
	PlexToken    string      `json:"plexToken"`
	PlexUsername string      `json:"plexUsername"`
	UserType     int         `json:"userType"`
	Permissions  int         `json:"permissions"`
	Avatar       string      `json:"avatar"`
	CreatedAt    string      `json:"createdAt"`
	UpdatedAt    string      `json:"updatedAt"`
	RequestCount int         `json:"requestCount"`
}

type UserSettings struct {
	Locale           string `json:"locale"`
	Region           string `json:"region"`
	OriginalLanguage string `json:"originalLanguage"`
}

type MovieResult struct {
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
	Title            string    `json:"title"`
	OriginalTitle    string    `json:"originalTitle"`
	ReleaseDate      string    `json:"releaseDate"`
	Adult            bool      `json:"adult"`
	Video            bool      `json:"video"`
	MediaInfo        MediaInfo `json:"mediaInfo"`
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

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Company struct {
	ID       int    `json:"id"`
	LogoPath string `json:"logoPath"`
	Name     string `json:"name"`
}

type ProductionCompany struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
}

type Network struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
}

type RelatedVideo struct {
	Url  string `json:"url"`
	Key  string `json:"key"`
	Name string `json:"name"`
	Size int    `json:"size"`
	Type string `json:"type"`
	Site string `json:"site"`
}

type MovieDetails struct {
	ID                  int                 `json:"id"`
	ImdbID              string              `json:"imdbId"`
	Adult               bool                `json:"adult"`
	BackdropPath        string              `json:"backdropPath"`
	PosterPath          string              `json:"posterPath"`
	Budget              int                 `json:"budget"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	RelatedVideos       []RelatedVideo      `json:"relatedVideos"`
	OriginalLanguage    string              `json:"originalLanguage"`
	OriginalTitle       string              `json:"originalTitle"`
	Overview            string              `json:"overview"`
	Popularity          float64             `json:"popularity"`
	ProductionCompanies []ProductionCompany `json:"productionCompanies"`
	ProductionCountries []ProductionCountry `json:"productionCountries"`
	ReleaseDate         string              `json:"releaseDate"`
	Releases            struct {
		Results []struct {
			Iso31661     string        `json:"iso_3166_1"`
			Rating       string        `json:"rating"`
			ReleaseDates []ReleaseDate `json:"release_dates"`
		} `json:"results"`
	} `json:"releases"`
	Revenue         int              `json:"revenue"`
	Runtime         int              `json:"runtime"`
	SpokenLanguages []SpokenLanguage `json:"spokenLanguages"`
	Status          string           `json:"status"`
	Tagline         string           `json:"tagline"`
	Title           string           `json:"title"`
	Video           bool             `json:"video"`
	VoteAverage     float64          `json:"voteAverage"`
	VoteCount       int              `json:"voteCount"`
	Credits         struct {
		Cast []Cast `json:"cast"`
		Crew []Crew `json:"crew"`
	} `json:"credits"`
	Collection     Collection      `json:"collection"`
	ExternalIds    ExternalIds     `json:"externalIds"`
	MediaInfo      MediaInfo       `json:"mediaInfo"`
	WatchProviders []WatchProvider `json:"watchProviders"`
}

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
	SpokenLanguages     []Language          `json:"spokenLanguages"`
	Seasons             []Season            `json:"seasons"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"voteAverage"`
	VoteCount           int                 `json:"voteCount"`
	Credits             Credits             `json:"credits"`
	ExternalIds         ExternalIds         `json:"externalIds"`
	Keywords            []Keyword           `json:"keywords"`
	MediaInfo           MediaInfo           `json:"mediaInfo"`
	WatchProviders      []WatchProvider     `json:"watchProviders"`
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
	Seasons   []struct {
		ID           int    `json:"id"`
		SeasonNumber int    `json:"seasonNumber"`
		Status       int    `json:"status"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
	} `json:"seasons"`
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
	FacebookID  interface{} `json:"facebookId"`
	FreebaseID  interface{} `json:"freebaseId"`
	FreebaseMid interface{} `json:"freebaseMid"`
	ImdbID      string      `json:"imdbId"`
	InstagramID interface{} `json:"instagramId"`
	TvdbID      int         `json:"tvdbId"`
	TvrageID    interface{} `json:"tvrageId"`
	TwitterID   interface{} `json:"twitterId"`
}

type PageInfo struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	Results int `json:"results"`

	// Not in schema
	PageSize int `json:"pageSize"`
}

type ContentRating struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Rating    string `json:"rating"`
}

type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Creator struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profilePath"`
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

type ProductionCountry struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

type Language struct {
	EnglishName string `json:"englishName"`
	Iso639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Credits struct {
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

type ReleaseDate struct {
	Certification string    `json:"certification"`
	Iso6391       string    `json:"iso_639_1"`
	Note          string    `json:"note"`
	ReleaseDate   time.Time `json:"release_date"`
	Type          int       `json:"type"`
}

type SpokenLanguage struct {
	EnglishName string `json:"englishName"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Collection struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"posterPath"`
	BackdropPath string `json:"backdropPath"`
}
