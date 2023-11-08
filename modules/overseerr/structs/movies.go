package overseerr_structs

import (
	"time"
)

type RelatedVideo struct {
	Url  string `json:"url"`
	Key  string `json:"key"`
	Name string `json:"name"`
	Size int    `json:"size"`
	Type string `json:"type"`
	Site string `json:"site"`
}


type ReleaseDate struct {
	Certification string    `json:"certification"`
	Iso6391       string    `json:"iso_639_1"`
	Note          string    `json:"note"`
	ReleaseDate   time.Time `json:"release_date"`
	Type          int       `json:"type"`
}

type Collection struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"posterPath"`
	BackdropPath string `json:"backdropPath"`
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
