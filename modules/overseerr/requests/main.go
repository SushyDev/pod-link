package overseerr_requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"time"
)

func GetMediaRequest(requestId int) (overseerr_structs.MediaRequest, error) {
	config := config.GetConfig()
	host := config.Settings.Overseerr.Host
	token := config.Settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/request/%v", host, requestId)

	if (config.Settings.Pod.Verbosity >= 2) {
		fmt.Printf("[DEBUG] %s\n", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return overseerr_structs.MediaRequest{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr_structs.MediaRequest{}, err
	}

	defer response.Body.Close()

	var data overseerr_structs.MediaRequest
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr_structs.MediaRequest{}, err
	}

	return data, nil
}

func GetPendingRequests() (RequestsReturned, error) {
	config := config.GetConfig()
	host := config.Settings.Overseerr.Host
	token := config.Settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/request?filter=processing&sort=added", host)

	if (config.Settings.Pod.Verbosity >= 2) {
		fmt.Printf("[DEBUG] %s\n", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return RequestsReturned{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return RequestsReturned{}, err
	}

	defer response.Body.Close()

	var requests RequestsReturned
	err = json.NewDecoder(response.Body).Decode(&requests)
	if err != nil {
		fmt.Println("Failed to decode response")
		return RequestsReturned{}, err
	}

	return requests, nil
}

func FilterRequests(requests []overseerr_structs.MediaRequest) ([]overseerr_structs.MediaRequest, error) {
	config := config.GetConfig()

	var filteredRequests []overseerr_structs.MediaRequest
	for _, request := range requests {
		date, err := time.Parse(time.RFC3339, request.CreatedAt)
		if err != nil {
			fmt.Println("Failed to parse date")
			return nil, err
		}

		currentDate := time.Now()
		difference := currentDate.Sub(date)

		if config.Settings.Pod.MissingContent.RequestAge == 0 {
			fmt.Println("Warning: Request age is set to 0, this will return all requests")
		}

		if difference.Hours() < config.Settings.Pod.MissingContent.RequestAge {
			continue
		}

		filteredRequests = append(filteredRequests, request)
	}

	return filteredRequests, nil
}

func GetRequestDetails(requestId int) (overseerr_structs.MediaRequest, error) {
	config := config.GetConfig()
	host := config.Settings.Overseerr.Host
	token := config.Settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/request/%v", host, requestId)
		
	if (config.Settings.Pod.Verbosity >= 2) {
		fmt.Printf("[DEBUG] %s\n", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return overseerr_structs.MediaRequest{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr_structs.MediaRequest{}, err
	}

	defer response.Body.Close()

	var request overseerr_structs.MediaRequest
	err = json.NewDecoder(response.Body).Decode(&request)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr_structs.MediaRequest{}, err
	}

	return request, nil
}
