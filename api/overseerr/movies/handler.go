package movies

import (
    "fmt"
    "os"
    "project-pod/api/debrid"
    "project-pod/api/plex"
    "project-pod/api/structs"
    "project-pod/api/torrentio"
    torrentio_movies "project-pod/api/torrentio/movies"
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
