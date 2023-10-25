package tv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pod-link/modules/torrentio"
	"regexp"
	"strings"
)

func GetList(ImdbId string, Season int, Episode int) []torrentio.Stream {
	realdebrid := os.Getenv("REAL_DEBRID_TOKEN")
	filter := "sort=qualitysize|qualityfilter=other,scr,cam,unknown|realdebrid=" + realdebrid
	url := fmt.Sprintf("https://torrentio.strem.fun/%s/stream/series/%s:%v:%v.json", filter, ImdbId, Season, Episode)

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
	for i, stream := range streams {
		isS := regexp.MustCompile(`(?i)[. ]s\d+[. ]`)
		isSeason := regexp.MustCompile(`(?i)[. ]season \d+[. ]`)
		isE := regexp.MustCompile(`(?i)[. ]e\d+[. ]`)
		isEpisode := regexp.MustCompile(`(?i)[. ]episode \d+[. ]`)

		if isE.MatchString(stream.Title) ||
			isEpisode.MatchString(stream.Title) ||
			!isS.MatchString(stream.Title) ||
			!isSeason.MatchString(stream.Title) {
			continue
		}

		results = append(results, streams[i])
	}

	return results
}

func FilterEpisodes(streams []torrentio.Stream) []torrentio.Stream {
	var results []torrentio.Stream
	for i, stream := range streams {
		url := strings.ReplaceAll(stream.Url, "https://torrentio.strem.fun/realdebrid/", "")
		url = strings.ReplaceAll(url, os.Getenv("REAL_DEBRID_TOKEN"), "")

		split := strings.Split(url, "/")
		if split[2] == "1" {
			results = append(results, streams[i])
		}
	}

	return results
}
