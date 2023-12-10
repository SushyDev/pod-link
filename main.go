package main

import (
	"fmt"
	"os"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_requests "pod-link/modules/overseerr/requests"
	overseerr_settings "pod-link/modules/overseerr/settings"
	overseerr_tv "pod-link/modules/overseerr/tv"
	"pod-link/modules/plex"
	"pod-link/modules/webhook"
	"time"
)

func missingContent() {
	requests, err := overseerr_requests.GetPendingRequests()
	if err != nil {
		fmt.Println(err)
		return
	}

	filteredRequests, err := overseerr_requests.FilterRequests(requests.Results)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, request := range filteredRequests {
		fmt.Printf("Received request with ID %d\n", request.ID)

		requestDetails, err := overseerr_requests.GetRequestDetails(request.ID)
		if err != nil {
			fmt.Println("Failed to get request details")
			fmt.Println(err)
			return
		}

		switch requestDetails.Media.MediaType {
		case "movie":
			overseerr_movies.Missing(requestDetails)
		case "tv":
			overseerr_tv.Missing(requestDetails)
		}

		time.Sleep(15 * time.Second)

		libraryIds := overseerr_settings.GetLibraryIdsByType(requestDetails.Media.MediaType)
		for _, libraryId := range libraryIds {
			plex.RefreshLibrary(libraryId)
		}
	}

	fmt.Println("Finished")
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
