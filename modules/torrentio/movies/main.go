package torrentio_movies

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pod-link/modules/config"
	"pod-link/modules/torrentio"
)

func GetStreams(ImdbId string) ([]torrentio.Stream, error) {
	baseURL := torrentio.GetBaseURL("movies")
	url := fmt.Sprintf("%s/stream/movie/%s.json", baseURL, ImdbId)

	if (config.GetConfig().Settings.Pod.Verbosity >= 2) {
		fmt.Printf("[DEBUG] %s\n", url)
	}

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

	switch response.StatusCode {
	case 200:
		var data torrentio.Response
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			fmt.Println("Failed to decode response")
			return nil, err
		}

		return data.Streams, nil
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to read response body")
		}

		fmt.Println(string(body))

		return nil, fmt.Errorf("Unknown error")
	}
}
