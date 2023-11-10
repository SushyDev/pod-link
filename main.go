package main

import (
	"fmt"
	"os"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_requests "pod-link/modules/overseerr/requests"
	overseerr_structs "pod-link/modules/overseerr/structs"
	overseerr_tv "pod-link/modules/overseerr/tv"
	"pod-link/modules/webhook"
	"sync"
)

func handleRequest(request overseerr_structs.MediaRequest, requestWg *sync.WaitGroup) {
	fmt.Printf("Received request with ID %d\n", request.ID)

	requestDetails, err := overseerr_requests.GetRequestDetails(request.ID)
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

	fmt.Printf("[%d] Done\n", request.ID)

	requestWg.Done()
}

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

	// perhaps split the requests into chunks and run each chunk concurrently
	var requestWg sync.WaitGroup
	for _, request := range filteredRequests {
		requestWg.Add(1)
		go handleRequest(request, &requestWg)
	}

	requestWg.Wait()

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
