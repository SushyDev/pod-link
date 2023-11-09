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

	fmt.Printf("[%v] %s\n", details.ImdbID, details.Title)

	streams, err := torrentio_movies.GetList(details.ImdbID)
	if err != nil {
		fmt.Println("Failed to get results")
		fmt.Println(err)
		return
	}

	streams = torrentio.FilterVersions(streams, "movies")

	if len(streams) == 0 {
		fmt.Println("No results found")
		return
	}

	for _, stream := range streams {
		properties, err := torrentio.GetPropertiesFromStream(stream)
		if err != nil {
			fmt.Println("Failed to get properties. Skipping")
			fmt.Println(err)
			continue
		}

		if !torrentio.MatchesProperties(stream, properties) {
			continue
		}

		fmt.Printf("[%s] %s\n", stream.Version, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Println("Failed to add magnet. Skipping")
			fmt.Println(err)
			continue
		}
	}
}
