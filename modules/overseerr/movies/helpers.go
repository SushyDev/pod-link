package overseerr_movies

import (
	"fmt"
	"pod-link/modules/debrid"
	"pod-link/modules/torrentio"
	torrentio_movies "pod-link/modules/torrentio/movies"
)

func FindById(movieId int) {
	details, err := GetMovieDetails(movieId)
	if err != nil {
		fmt.Println("Failed to get details")
		fmt.Println(err)
		return
	}

	if details.ImdbID == "" {
		fmt.Printf("[%v] No IMDB ID found\n", movieId)
		return
	}

	fmt.Printf("[%v] %s\n", movieId, details.Title)

	streams, err := torrentio_movies.GetStreams(details.ImdbID)
	if err != nil {
		fmt.Printf("[%v] Failed to get streams\n", movieId)
		fmt.Println(err)
		return
	}

	streams = torrentio.FilterVersions(streams, "movies")

	if len(streams) == 0 {
		fmt.Printf("[%v] No matching streams found\n", movieId)
		return
	}

	for _, stream := range streams {
		properties, err := torrentio.GetPropertiesFromStream(stream)
		if err != nil {
			fmt.Printf("[%v] Failed to get properties\n", movieId)
			fmt.Println(err)
			continue
		}

		if !torrentio.MatchesProperties(stream, properties) {
			continue
		}

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Printf("[%v] Failed to add magnet\n", movieId)
			fmt.Println(err)
			continue
		}

		fmt.Printf("[%s] + %s\n", stream.Version, properties.Title)
	}
}
