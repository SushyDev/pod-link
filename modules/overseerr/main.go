package overserr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr "pod-link/modules/overseerr/structs"
	overseerr_tv "pod-link/modules/overseerr/tv"
	"pod-link/modules/structs"
	"time"
)

func HandleMediaAutoApprovedNotification(notification structs.MediaAutoApprovedNotification) {
	switch notification.Media.MediaType {
	case "movie":
		overseerr_movies.Request(notification)
	case "tv":
		overseerr_tv.Request(notification)
	}
}

type RequestsReturned struct {
	PageInfo overseerr.PageInfo       `json:"pageInfo"`
	Results  []overseerr.MediaRequest `json:"results"`
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

func FilterRequests(requests []overseerr.MediaRequest) ([]overseerr.MediaRequest, error) {
	var filteredRequests []overseerr.MediaRequest
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

func GetRequestDetails(requestId int) (overseerr.MediaRequest, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/request/%v", host, requestId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return overseerr.MediaRequest{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr.MediaRequest{}, err
	}

	defer response.Body.Close()

	var request overseerr.MediaRequest
	err = json.NewDecoder(response.Body).Decode(&request)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr.MediaRequest{}, err
	}

	return request, nil
}

type Season struct {
    ID int `json:"id"`
    SeasonNumber int `json:"seasonNumber"`
    Status int `json:"status"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

func FilterCompleteSeasons(details overseerr.MediaRequest) []int {
    var seasons []int
    for _, season := range details.Seasons {
        for _, mediaInfoSeason := range details.Media.Seasons {
            if season.SeasonNumber == mediaInfoSeason.SeasonNumber && mediaInfoSeason.Status != 5 {
                seasons = append(seasons, mediaInfoSeason.SeasonNumber)
            }
        }
    }

    return seasons
}
