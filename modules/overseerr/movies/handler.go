package movies

import (
	"fmt"
	"os"
	"pod-link/modules/debrid"
	"pod-link/modules/plex"
	"pod-link/modules/structs"
	"pod-link/modules/torrentio"
	torrentio_movies "pod-link/modules/torrentio/movies"
)

func Request(notification structs.MediaAutoApprovedNotification) {
    details := GetDetails(notification.Media.TmdbId)

    results := torrentio_movies.GetList(details.ExternalIds.ImdbId)
    for _, result := range results {
        properties := torrentio.GetPropertiesFromStream(result)

        fmt.Println("Adding:", properties.Title)
        fmt.Println("    -", properties.Link)

        err := debrid.AddMagnet(properties.Link, properties.Files)
        if err != nil {
            fmt.Println("\033[31m", err, "\033[0m")
        }
    }

    err := plex.RefreshLibrary(os.Getenv("PLEX_MOVIE_ID"))
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to refresh library")
    }
}
