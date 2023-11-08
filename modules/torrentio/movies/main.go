package torrentio_movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/torrentio"
)

func GetList(ImdbId string) ([]torrentio.Stream, error) {
	baseURL := torrentio.GetBaseURL("movies")
	url := fmt.Sprintf("%s/stream/movie/%s.json", baseURL, ImdbId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return nil, err
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return nil, err
	}

	defer response.Body.Close()

	var data torrentio.Response
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return nil, err
	}

	return data.Streams, nil
}
