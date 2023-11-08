package main

import (
	"fmt"
	"os"
	"pod-link/modules/overseerr"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"pod-link/modules/plex"
	"pod-link/modules/webhook"
)

func getDirectoryBySeason(directories ([]plex.TvDirectory), season int) (plex.TvDirectory, error) {
	for _, directory := range directories {
		if directory.Index == fmt.Sprintf("%v", season) {
			return directory, nil
		}
	}

	return plex.TvDirectory{}, fmt.Errorf("No directory found for season %v", season)
}

func missingTvContent(requestDetails overseerr_structs.MediaRequest) {
	seasons := overseerr.FilterCompleteSeasons(requestDetails)

	if len(seasons) == 0 {
		fmt.Printf("[%v] No incomplete seasons\n", requestDetails.ID)
		return
	}

	// if no rating key then content is not on plex so follow normal content add flow
	if requestDetails.Media.RatingKey == "" {
		fmt.Printf("[%v] No rating key\n", requestDetails.ID)
		// todo
		return
	}

	tvMetadata, err := plex.GetTvMetadata(requestDetails.Media.RatingKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, season := range seasons {
		directory, err := getDirectoryBySeason(tvMetadata.Directory, season)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("[%v] %s\n", requestDetails.ID, directory.Title)

		plex.GetSeasonMetadata(directory.RatingKey)
	}

	// overseerr_tv.FindById(requestDetails.Media.TmdbID, seasons)
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
		fmt.Println(request.ID)
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
