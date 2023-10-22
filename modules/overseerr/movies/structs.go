package movies

import (
    "time"
)

type Genre struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

type Video struct {
    Url  string `json:"url"`
    Key  string `json:"key"`
    Name string `json:"name"`
    Size int    `json:"size"`
    Type string `json:"type"`
    Site string `json:"site"`
}

type ProductionCompany struct {
    Id            int    `json:"id"`
    LogoPath      string `json:"logoPath"`
    OriginCountry string `json:"originCountry"`
    Name          string `json:"name"`
}

type ProductionCountry struct {
    Iso31661 string `json:"iso_3166_1"`
    Name     string `json:"name"`
}

type ReleaseDate struct {
    Certification string    `json:"certification"`
    Iso6391       string    `json:"iso_639_1"`
    Note          string    `json:"note"`
    ReleaseDate   time.Time `json:"release_date"`
    Type          int       `json:"type"`
}

type Cast struct {
    Id        int    `json:"id"`
    CastId    int    `json:"castId"`
    Character string `json:"character"`
    CreditId  string `json:"creditId"`
    Gender    int    `json:"gender"`
    Name      string `json:"name"`
    Order     int    `json:"order"`
    ProfilePath string `json:"profilePath"`
}

type Crew struct {
    Id        int    `json:"id"`
    CreditId  string `json:"creditId"`
    Gender    int    `json:"gender"`
    Name      string `json:"name"`
    Job       string `json:"job"`
    Department string `json:"department"`
    ProfilePath string `json:"profilePath"`
}

type Collection struct {
    Id          int    `json:"id"`
    Name        string `json:"name"`
    PosterPath  string `json:"posterPath"`
    BackdropPath string `json:"backdropPath"`
}

type ExternalIds struct {
    FacebookId string `json:"facebookId"`
    FreebaseId string `json:"freebaseId"`
    FreebaseMid string `json:"freebaseMid"`
    ImdbId     string `json:"imdbId"`
    InstagramId string `json:"instagramId"`
    TvdbId     int    `json:"tvdbId"`
    TvrageId   int    `json:"tvrageId"`
    TwitterId  string `json:"twitterId"`
}

type RequestedBy struct {
    Id          int    `json:"id"`
    Email       string `json:"email"`
    Username    string `json:"username"`
    PlexToken   string `json:"plexToken"`
    PlexUsername string `json:"plexUsername"`
    UserType    int    `json:"userType"`
    Permissions int    `json:"permissions"`
    Avatar      string `json:"avatar"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
    RequestCount int  `json:"requestCount"`
}

type WatchProvider struct {
    Iso31661 string `json:"iso_3166_1"`
    Link     string `json:"link"`
    Buy      []struct {
        DisplayPriority int    `json:"displayPriority"`
        LogoPath        string `json:"logoPath"`
        Id              int    `json:"id"`
        Name            string `json:"name"`
    } `json:"buy"`
    Flatrate []struct {
        DisplayPriority int    `json:"displayPriority"`
        LogoPath        string `json:"logoPath"`
        Id              int    `json:"id"`
        Name            string `json:"name"`
    } `json:"flatrate"`
}

type Movie struct {
    Id                int               `json:"id"`
    ImdbId            string            `json:"imdbId"`
    Adult             bool              `json:"adult"`
    BackdropPath      string            `json:"backdropPath"`
    PosterPath        string            `json:"posterPath"`
    Budget            int               `json:"budget"`
    Genres            []Genre           `json:"genres"`
    Homepage          string            `json:"homepage"`
    RelatedVideos     []Video           `json:"relatedVideos"`
    OriginalLanguage  string            `json:"originalLanguage"`
    OriginalTitle     string            `json:"originalTitle"`
    Overview           string           `json:"overview"`
    Popularity         float64           `json:"popularity"`
    ProductionCompanies []ProductionCompany `json:"productionCompanies"`
    ProductionCountries []ProductionCountry `json:"productionCountries"`
    ReleaseDate        string             `json:"releaseDate"`
    Releases           struct {
        Results []struct {
            Iso31661     string        `json:"iso_3166_1"`
            Rating       string        `json:"rating"`
            ReleaseDates []ReleaseDate `json:"release_dates"`
        } `json:"results"`
    } `json:"releases"`
    Revenue         int    `json:"revenue"`
    Runtime         int    `json:"runtime"`
    SpokenLanguages []struct {
        EnglishName string `json:"englishName"`
        Iso6391     string `json:"iso_639_1"`
        Name        string `json:"name"`
    } `json:"spokenLanguages"`
    Status   string `json:"status"`
    Tagline  string `json:"tagline"`
    Title    string `json:"title"`
    Video    bool   `json:"video"`
    VoteAverage float64 `json:"voteAverage"`
    VoteCount   int    `json:"voteCount"`
    Credits     struct {
        Cast []Cast `json:"cast"`
        Crew []Crew `json:"crew"`
    } `json:"credits"`
    Collection   Collection `json:"collection"`
    ExternalIds  ExternalIds `json:"externalIds"`
    MediaInfo   struct {
        ID        int              `json:"id"`
        TmdbID    int              `json:"tmdbId"`
        TvdbID    int              `json:"tvdbId"`
        Status    int              `json:"status"`
        CreatedAt string           `json:"createdAt"`
        UpdatedAt string           `json:"updatedAt"`
    } `json:"mediaInfo"`
    WatchProviders []struct {
        WatchProvider
    } `json:"watchProviders"`
}
