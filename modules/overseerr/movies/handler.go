package movies

import (
	"fmt"
	"os"
	"link-pod/modules/debrid"
	"link-pod/modules/plex"
	"link-pod/modules/structs"
	"link-pod/modules/torrentio"
	torrentio_movies "link-pod/modules/torrentio/movies"
)

func Request(notification structs.MediaAutoApprovedNotification) {
    details := GetDetails(notification.Media.TmdbId)

    results := torrentio_movies.GetList(details.ExternalIds.ImdbId)
    for _, result := range results {
        properties := torrentio.GetPropertiesFromStream(result)

        fmt.Println("Adding:", properties.Title)

        err := debrid.AddMagnet(properties.Link)
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
