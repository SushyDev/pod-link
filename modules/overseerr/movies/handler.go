package movies

import (
	"fmt"
	"pod-link/modules/config"
	"pod-link/modules/debrid"
	"pod-link/modules/plex"
	"pod-link/modules/structs"
	"pod-link/modules/torrentio"
	torrentio_movies "pod-link/modules/torrentio/movies"
	"strings"
	"time"
)

func FilterProperties(results []torrentio.Stream) []torrentio.Stream {
	var filtered []torrentio.Stream

	config := config.GetConfig()

	for _, result := range results {
		properties, err := torrentio.GetPropertiesFromStream(result)
		if err != nil {
			fmt.Println("Failed to get properties. Skipping.")
			fmt.Println(err)
			continue
		}

		if properties.Files == "all" {
			filtered = append(filtered, result)
			continue
		}

		fileCount := len(strings.Split(properties.Files, ","))
		maxFiles := config.Movies.MaxFiles

		if fileCount <= maxFiles {
			filtered = append(filtered, result)
		}
	}

	return filtered
}

func Request(notification structs.MediaAutoApprovedNotification) {
	details, err := GetDetails(notification.Media.TmdbId)
	if err != nil {
		fmt.Println("Failed to get details")
		fmt.Println(err)
		return
	}

	fmt.Println("Got request for", details.Title)

	results, err := torrentio_movies.GetList(details.ExternalIds.ImdbId)
	if err != nil {
		fmt.Println("Failed to get results")
		fmt.Println(err)
		return
	}

	results = torrentio.FilterVersions(results, "movies")
	results = FilterProperties(results)

	if len(results) == 0 {
		fmt.Println("No results found")
		return
	}

	for _, result := range results {
		properties, err := torrentio.GetPropertiesFromStream(result)
		if err != nil {
			fmt.Println("Failed to get properties. Skipping")
			fmt.Println(err)
			continue
		}

		fmt.Printf("[%s] %s\n", result.Version, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Println("Failed to add magnet. Skipping")
			fmt.Println(err)
			continue
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
			fmt.Println("Failed to refresh library")
			fmt.Println(err)
		}
	}
}
