package overseerr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	overseerr_settings "pod-link/modules/overseerr/settings"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"time"
)

func GetMediaRequest(requestId int) (overseerr_structs.MediaRequest, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/request/%v", host, requestId)

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

type RequestsReturned struct {
	PageInfo overseerr_structs.PageInfo       `json:"pageInfo"`
	Results  []overseerr_structs.MediaRequest `json:"results"`
}

func GetPendingRequests() (RequestsReturned, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/request?filter=processing&sort=added", host)

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
	var filteredRequests []overseerr_structs.MediaRequest
	for _, request := range requests {
		date, err := time.Parse(time.RFC3339, request.CreatedAt)
		if err != nil {
			fmt.Println("Failed to parse date")
			return nil, err
		}

		currentDate := time.Now()
		difference := currentDate.Sub(date)

		if difference.Hours() > 24 {
			filteredRequests = append(filteredRequests, request)
		}
	}

	return filteredRequests, nil
}

func GetRequestDetails(requestId int) (overseerr_structs.MediaRequest, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/request/%v", host, requestId)

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

func getServerConnection(connections []overseerr_structs.PlexConnection) (overseerr_structs.PlexConnection, error) {
	for _, connection := range connections {
		if connection.Status == 200 {
			return connection, nil
		}
	}

	return overseerr_structs.PlexConnection{}, nil
}

func GetPlexTokenAndHost() (string, string, error) {
	plexSettings, err := overseerr_settings.GetPlexSettings()
	if err != nil {
		return "", "", err
	}

	machineId := plexSettings.MachineID

	plexServers, err := overseerr_settings.GetPlexServers()
	if err != nil {
		return "", "", err
	}

	for _, server := range plexServers {
		if server.ClientIdentifier == machineId {
			connection, err := getServerConnection(server.Connection)
			if err != nil {
				return "", "", err
			}

			return server.AccessToken, connection.Uri, nil
		}
	}

	return "", "", nil
}
