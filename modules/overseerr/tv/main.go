package overseerr_tv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"strconv"
	"strings"
)

func GetTvDetails(tvId int) (overseerr_structs.TvDetails, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token
	url := fmt.Sprintf("%s/api/v1/tv/%v", host, tvId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return overseerr_structs.TvDetails{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr_structs.TvDetails{}, err
	}

	defer response.Body.Close()

	var data overseerr_structs.TvDetails
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr_structs.TvDetails{}, err
	}

	return data, nil
}

func GetSeasonDetails(tvId int, seasonId int) (overseerr_structs.Season, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/tv/%v/season/%v", host, tvId, seasonId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return overseerr_structs.Season{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr_structs.Season{}, err
	}

	defer response.Body.Close()

	var data overseerr_structs.Season
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr_structs.Season{}, err
	}

	return data, nil
}

type OverseerrResponse struct {
	ExternalIds struct {
		ImdbId string `json:"imdbId"`
	} `json:"externalIds"`
}

func isAnime(keywords []overseerr_structs.Keyword) bool {
	for _, keyword := range keywords {
		if strings.ToLower(keyword.Name) == "anime" {
			return true
		}
	}

	return false
}

func getEpisodeCountBySeason(tvId int, seasonId int) ([]int, error) {
	details, err := GetSeasonDetails(tvId, seasonId)
	if err != nil {
		fmt.Println("Failed to get season details")
		return nil, err
	}

	var episodeNumbers = []int{}
	for _, episode := range details.Episodes {
		if episode.AirDate == "" {
			continue
		}

		episodeNumbers = append(episodeNumbers, episode.EpisodeNumber)
	}

	return episodeNumbers, nil
}

func getRequestedSeasons(extra []overseerr_structs.Extra) []int {
	var seasonNumbers = []int{}

	for _, extra := range extra {
		if extra.Name != "Requested Seasons" {
			continue
		}

		list := strings.Split(extra.Value, ", ")
		for _, season := range list {
			seasonNumber, err := strconv.Atoi(season)
			if err != nil {
				fmt.Println("Failed to convert season to int. Skipping")
				fmt.Println(err)
				continue
			}

			seasonNumbers = append(seasonNumbers, seasonNumber)
		}
	}

	return seasonNumbers
}
