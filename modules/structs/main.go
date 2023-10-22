package structs

type Extra struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

type MediaAutoApprovedNotification struct {
    NotificationType string `json:"notification_type"`
    Media            struct {
        MediaType string `json:"media_type"`
        TmdbId    string `json:"tmdbId"`
        TvdId     string `json:"tvdbId"`
    }
    Extra []Extra `json:"extra"`
}
