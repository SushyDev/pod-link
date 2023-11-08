package overseerr_structs

type MediaAutoApprovedNotification struct {
	NotificationType string `json:"notification_type"`
	Media            struct {
		MediaType string `json:"media_type"`
		TmdbId    string `json:"tmdbId"`
		TvdId     string `json:"tvdbId"`
	}
	Extra []Extra `json:"extra"`
}

type UserSettings struct {
	Locale           string `json:"locale"`
	Region           string `json:"region"`
	OriginalLanguage string `json:"originalLanguage"`
}

type Company struct {
	ID       int    `json:"id"`
	LogoPath string `json:"logoPath"`
	Name     string `json:"name"`
}

type PageInfo struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	Results int `json:"results"`

	// Not in schema
	PageSize int `json:"pageSize"`
}

type Language struct {
	EnglishName string `json:"englishName"`
	Iso639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Extra struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
