package plex

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"pod-link/modules/overseerr"
)

func RefreshLibrary(id string) error {
	token, host, err := overseerr.GetPlexTokenAndHost()
	if err != nil {
		fmt.Println("Failed to get plex token and host")
		return err
	}

	url := fmt.Sprintf("%s/library/sections/%s/refresh?X-Plex-Token=%s", host, id, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create request")
		return err
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to send request")
		return err
	}

	defer response.Body.Close()

	return nil
}

func GetTvMetadata(ratingKey string) (TvMetadata, error) {
	token, host, err := overseerr.GetPlexTokenAndHost()
	if err != nil {
		fmt.Println("Failed to get plex token and host")
		return TvMetadata{}, err
	}

	url := fmt.Sprintf("%s/library/metadata/%v/children?X-Plex-Token=%s", host, ratingKey, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return TvMetadata{}, err
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return TvMetadata{}, err
	}

	defer response.Body.Close()

	var data TvMetadata
	err = xml.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return TvMetadata{}, err
	}

	return data, nil
}

func GetSeasonMetadata(ratingKey string) {
	token, host, err := overseerr.GetPlexTokenAndHost()
	if err != nil {
		fmt.Println("Failed to get plex token and host")
		return
	}

	url := fmt.Sprintf("%s/library/metadata/%v/children?X-Plex-Token=%s", host, ratingKey, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create request")
		return
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to send request")
		return
	}

	defer response.Body.Close()

	var data SeasonMetadata
	err = xml.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to decode response")
		return
	}


	for _, episode := range data.Video {
		fmt.Printf("[%v] %s\n", episode.RatingKey, episode.Title)
	}
}
