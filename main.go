package main

import (
	"fmt"
	"os"
	overseerr "pod-link/modules/overseerr"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_structs "pod-link/modules/overseerr/structs"
	overseerr_tv "pod-link/modules/overseerr/tv"
	"pod-link/modules/webhook"
)

func missingTvContent(requestDetails overseerr_structs.MediaRequest) {
	seasons := overseerr.FilterCompleteSeasons(requestDetails)

	if len(seasons) == 0 {
		fmt.Printf("[%v] No incomplete seasons\n", requestDetails.ID)
		return
	}

	overseerr_tv.FindById(requestDetails.Media.TmdbID, seasons)
}

func missingMovieContent(requestDetails overseerr_structs.MediaRequest) {
	overseerr_movies.FindById(requestDetails.Media.TmdbID)
}

func missingContent() {
	requests, err := overseerr.GetPendingRequests()
	if err != nil {
		fmt.Println(err)
		return
	}

	// filteredRequests, err := overseerr.FilterRequests(requests.Results)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	for _, request := range requests.Results {
		requestDetails, err := overseerr.GetRequestDetails(request.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch(requestDetails.Media.MediaType) {
		case "movie":
			missingMovieContent(requestDetails)
		case "tv":
			missingTvContent(requestDetails)
		}

	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		webhook.Listen()
		return
	}

	switch args[0] {
	case "missing-content":
		missingContent()
	default:
		err := fmt.Errorf("Unknown command")
		fmt.Println(err)
	}
}
