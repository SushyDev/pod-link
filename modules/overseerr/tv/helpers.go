package overseerr_tv

import (
	"fmt"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"pod-link/modules/plex"
	"strconv"
	"strings"
	"time"
)

func episodeIsStored(episode int, storedEpisodes ([]plex.Video)) bool {
	for _, storedEpisode := range storedEpisodes {
		if storedEpisode.Index == fmt.Sprintf("%d", episode) {
			return true
		}
	}

	return false
}

func filterCompleteSeasons(details overseerr_structs.MediaRequest) []int {
	var seasons []int
	for _, season := range details.Seasons {
		for _, mediaInfoSeason := range details.Media.Seasons {
			if season.SeasonNumber == mediaInfoSeason.SeasonNumber && mediaInfoSeason.Status == 5 {
				continue
			}
		}

		seasons = append(seasons, season.SeasonNumber)
	}

	return seasons
}

func isAnime(keywords []overseerr_structs.Keyword) bool {
	for _, keyword := range keywords {
		if strings.ToLower(keyword.Name) == "anime" {
			return true
		}
	}

	return false
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

func getEpisodeNumbersBySeason(details overseerr_structs.Season, seasonID int) []int {
	var episodeNumbers = []int{}
	for _, episode := range details.Episodes {
		episodeNumbers = append(episodeNumbers, episode.EpisodeNumber)
	}

	return episodeNumbers
}

func getReleasedEpisodeNumbersBySeason(details overseerr_structs.Season, seasonID int) []int {
	var episodeNumbers = []int{}
	for _, episode := range details.Episodes {
		if episode.AirDate == "" {
			continue
		}

		if episode.AirDate > time.Now().Format("2006-01-02") {
			continue
		}

		episodeNumbers = append(episodeNumbers, episode.EpisodeNumber)
	}

	return episodeNumbers
}
