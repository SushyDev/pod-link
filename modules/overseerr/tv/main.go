package tv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	"pod-link/modules/structs"
	"strconv"
	"strings"
)

type OverseerrResponse struct {
	ExternalIds struct {
		ImdbId string `json:"imdbId"`
	} `json:"externalIds"`
}

func isAnime(keywords []Keyword) bool {
	for _, keyword := range keywords {
		if strings.ToLower(keyword.Name) == "anime" {
			return true
		}
	}

	return false
}

func getEpisodeCountBySeason(number int, seasons []Season) int {
	for _, season := range seasons {
		if season.SeasonNumber == number {
			return season.EpisodeCount
		}
	}

	return 0
}

func getRequestedSeasons(extra []structs.Extra) []int {
	var seasonNumbers = []int{}

	for _, extra := range extra {
		if extra.Name != "Requested Seasons" {
			continue
		}

		list := strings.Split(extra.Value, ", ")
		for _, season := range list {
			seasonNumber, err := strconv.Atoi(season)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Failed to convert season to int")
			}

			seasonNumbers = append(seasonNumbers, seasonNumber)
		}
	}

	return seasonNumbers
}

func GetDetails(id string) Tv {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token
	url := fmt.Sprintf("%s/api/v1/tv/%s", host, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create request")
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to send request")
	}

	defer response.Body.Close()

	var data Tv
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to decode response")
	}

	return data
}
