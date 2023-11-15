package plex

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"pod-link/modules/config"
	"pod-link/modules/overseerr"
)

func RefreshLibrary(id string) error {
	token, host, err := overseerr.GetPlexTokenAndHost()
	if err != nil {
		fmt.Println("Failed to get plex token and host")
		return err
	}

	url := fmt.Sprintf("%s/library/sections/%s/refresh?X-Plex-Token=%s", host, id, token)

	if (config.GetConfig().Settings.Pod.Verbosity >= 2) {
		fmt.Printf("[DEBUG] %s", url)
	}

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

func GetEpisodesBySeason(video []Video, season int) []Video {
	var episodes []Video
	for _, episode := range video {
		if episode.ParentIndex == fmt.Sprintf("%v", season) {
			episodes = append(episodes, episode)
		}
	}

	return episodes
}

func GetShowLeaves(ratingKey string) (ShowLeaves, error) {
	token, host, err := overseerr.GetPlexTokenAndHost()
	if err != nil {
		fmt.Println("Failed to get plex token and host")
		return ShowLeaves{}, err
	}

	url := fmt.Sprintf("%s/library/metadata/%v/allLeaves?X-Plex-Token=%s", host, ratingKey, token)
	
	if (config.GetConfig().Settings.Pod.Verbosity >= 2) {
		fmt.Printf("[DEBUG] %s", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return ShowLeaves{}, err
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return ShowLeaves{}, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case 200:
		var data ShowLeaves
		err = xml.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			fmt.Println("Failed to decode response")
			return ShowLeaves{}, err
		}

		return data, nil
	case 404:
		return ShowLeaves{}, fmt.Errorf("[%v] Could not find show on Plex", response.StatusCode)
	default:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to read response body")
			return ShowLeaves{}, err
		}

		fmt.Println(string(body))

		return ShowLeaves{}, fmt.Errorf("[%v] Unknown error", response.StatusCode)
	}
}
