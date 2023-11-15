package overseerr_movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	overseerr_structs "pod-link/modules/overseerr/structs"
)

func GetMovieDetails(movieId int) (overseerr_structs.MovieDetails, error) {
	config := config.GetConfig()
	host := config.Settings.Overseerr.Host
	token := config.Settings.Overseerr.Token
	url := fmt.Sprintf("%s/api/v1/movie/%v", host, movieId)

	if (config.Settings.Pod.Verbosity >= 2) {
		fmt.Printf("[DEBUG] %s\n", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return overseerr_structs.MovieDetails{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr_structs.MovieDetails{}, err
	}

	defer response.Body.Close()

	var data overseerr_structs.MovieDetails
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr_structs.MovieDetails{}, err
	}

	return data, nil
}
