package overseerr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	overseerr "pod-link/modules/overseerr/structs"
)

func GetMovieDetails(movieId int) (overseerr.MovieDetails, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token
	url := fmt.Sprintf("%s/api/v1/movie/%v", host, movieId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return overseerr.MovieDetails{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr.MovieDetails{}, err
	}

	defer response.Body.Close()

	var data overseerr.MovieDetails
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr.MovieDetails{}, err
	}

	return data, nil
}

func GetTvDetails(tvId int) (overseerr.TvDetails, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token
	url := fmt.Sprintf("%s/api/v1/tv/%v", host, tvId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return overseerr.TvDetails{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr.TvDetails{}, err
	}

	defer response.Body.Close()

	var data overseerr.TvDetails
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr.TvDetails{}, err
	}

	return data, nil
}

func GetSeasonDetails(tvId int, seasonId int) (overseerr.Season, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/tv/%v/season/%v", host, tvId, seasonId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return overseerr.Season{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr.Season{}, err
	}

	defer response.Body.Close()

	var data overseerr.Season
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr.Season{}, err
	}

	return data, nil
}

func GetMediaRequest(requestId int) (overseerr.MediaRequest, error) {
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

	var data overseerr.MediaRequest
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr.MediaRequest{}, err
	}

	return data, nil
}
