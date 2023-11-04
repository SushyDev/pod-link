package tv

import (
	"fmt"
	overseerr_api "pod-link/modules/overseerr/api"
	overseerr "pod-link/modules/overseerr/structs"
	"pod-link/modules/structs"
	"strconv"
	"strings"
)

type OverseerrResponse struct {
	ExternalIds struct {
		ImdbId string `json:"imdbId"`
	} `json:"externalIds"`
}

func isAnime(keywords []overseerr.Keyword) bool {
	for _, keyword := range keywords {
		if strings.ToLower(keyword.Name) == "anime" {
			return true
		}
	}

	return false
}

func getEpisodeCountBySeason(tvId int, seasonId int) ([]int, error) {
	details, err := overseerr_api.GetSeasonDetails(tvId, seasonId)
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
				fmt.Println("Failed to convert season to int. Skipping")
				fmt.Println(err)
				continue
			}

			seasonNumbers = append(seasonNumbers, seasonNumber)
		}
	}

	return seasonNumbers
}
