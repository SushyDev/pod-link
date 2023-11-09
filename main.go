package main

import (
	"fmt"
	"os"
	"pod-link/modules/overseerr"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_structs "pod-link/modules/overseerr/structs"
	overseerr_tv "pod-link/modules/overseerr/tv"
	"pod-link/modules/webhook"
	"sync"
)

func handleRequest(request overseerr_structs.MediaRequest, requestWg *sync.WaitGroup) {
	fmt.Printf("Received request with ID %d\n", request.ID)

	requestDetails, err := overseerr.GetRequestDetails(request.ID)
	if err != nil {
		fmt.Println("Failed to get request details. Skipping")
		fmt.Println(err)
		return
	}

	switch requestDetails.Media.MediaType {
	case "movie":
		overseerr_movies.Missing(requestDetails)
	case "tv":
		overseerr_tv.Missing(requestDetails)
	}

	requestWg.Done()
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

	var requestWg sync.WaitGroup
	for _, request := range requests.Results {
		requestWg.Add(1)
		handleRequest(request, &requestWg)
	}

	requestWg.Wait()
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
