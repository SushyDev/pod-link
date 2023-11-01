package tv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	"pod-link/modules/torrentio"
	"regexp"
	"strings"
)

func GetList(ImdbId string, Season int, Episode int) []torrentio.Stream {
	baseURL := torrentio.GetBaseURL()
	url := fmt.Sprintf("%s/stream/series/%s:%v:%v.json", baseURL, ImdbId, Season, Episode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create request")
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to send request")
	}

	defer response.Body.Close()

	var data torrentio.Response
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to decode response")
	}

	return data.Streams
}

func FilterSeasons(streams []torrentio.Stream) []torrentio.Stream {
	var results []torrentio.Stream

	config := config.GetConfig()

	for i, stream := range streams {
		var isSeasonMatch bool
		var isEpisodeMatch bool

		for _, season := range config.Shows.Seasons {
			isSeasonMatch = regexp.MustCompile(season).MatchString(stream.Title)

			if isSeasonMatch {
				break
			}
		}

		for _, episode := range config.Shows.Episodes {
			isEpisodeMatch = regexp.MustCompile(episode).MatchString(stream.Title)

			if isEpisodeMatch {
				break
			}
		}

		if isSeasonMatch && !isEpisodeMatch {
			results = append(results, streams[i])
		}
	}

	return results
}

func FilterEpisodes(streams []torrentio.Stream) []torrentio.Stream {
	var results []torrentio.Stream
	for i, stream := range streams {
		url := strings.ReplaceAll(stream.Url, "https://torrentio.strem.fun/realdebrid/", "")
		settings := config.GetSettings()
		token := settings.RealDebrid.Token

		url = strings.ReplaceAll(url, token, "")

		split := strings.Split(url, "/")
		if split[2] == "1" {
			results = append(results, streams[i])
		}
	}

	return results
}
