package torrentio_tv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	"pod-link/modules/torrentio"
	"regexp"
	"strings"
)

func GetList(ImdbId string, Season int, Episode int) ([]torrentio.Stream, error) {
	baseURL := torrentio.GetBaseURL("shows")
	url := fmt.Sprintf("%s/stream/series/%s:%v:%v.json", baseURL, ImdbId, Season, Episode)

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

func isSeasonOrEpisodeMatch(stream torrentio.Stream) (bool, bool, error) {
	var isSeasonMatch bool
	var isEpisodeMatch bool

	config := config.GetConfig()

	streamTitle := strings.Split(stream.Title, "\n")[0]

	for _, season := range config.Shows.Seasons {
		regex, err := regexp.Compile(season)
		if err != nil {
			fmt.Println("Error compiling regular expression")
			return false, false, err
		}

		isSeasonMatch = regex.MatchString(streamTitle)

		if isSeasonMatch {
			break
		}
	}

	for _, episode := range config.Shows.Episodes {
		regex, err := regexp.Compile(episode)
		if err != nil {
			fmt.Println("Error compiling regular expression")
			return false, false, err
		}

		isEpisodeMatch = regex.MatchString(streamTitle)

		if isEpisodeMatch {
			break
		}
	}

	return isSeasonMatch, isEpisodeMatch, nil
}


func FilterSeasons(streams []torrentio.Stream) ([]torrentio.Stream, error) {
	var results []torrentio.Stream

	for i, stream := range streams {
		isSeasonMatch, isEpisodeMatch, err := isSeasonOrEpisodeMatch(stream)
		if err != nil {
			return nil, err
		}

		if isSeasonMatch && !isEpisodeMatch {
			results = append(results, streams[i])
		}
	}

	return results, nil
}

func FilterEpisodes(streams []torrentio.Stream) []torrentio.Stream {
	var results []torrentio.Stream

	for i, stream := range streams {
		_, isEpisodeMatch, err := isSeasonOrEpisodeMatch(stream)
		if err != nil {
			return nil
		}

		if isEpisodeMatch {
			results = append(results, streams[i])
		}
	}

	return results
}
