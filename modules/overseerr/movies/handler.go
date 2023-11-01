package movies

import (
	"fmt"
	"pod-link/modules/config"
	"pod-link/modules/debrid"
	"pod-link/modules/plex"
	"pod-link/modules/structs"
	"pod-link/modules/torrentio"
	torrentio_movies "pod-link/modules/torrentio/movies"
	"time"
)

func Request(notification structs.MediaAutoApprovedNotification) {
	details := GetDetails(notification.Media.TmdbId)

	results := torrentio_movies.GetList(details.ExternalIds.ImdbId)
	for _, result := range results {
		properties := torrentio.GetPropertiesFromStream(result)

		fmt.Println("Adding:", properties.Title)

		err := debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Println("\033[31m", err, "\033[0m")
		}
	}

	settings := config.GetSettings()
	host := settings.Plex.Host
	token := settings.Plex.Token
	movieId := settings.Plex.MovieId

	if host != "" && token != "" && movieId != "" {
		time.Sleep(20 * time.Second)
		err := plex.RefreshLibrary(movieId)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to refresh library")
		}
	}
}
