package overseerr_tv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	"pod-link/modules/debrid"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"pod-link/modules/plex"
	"pod-link/modules/torrentio"
	torrentio_tv "pod-link/modules/torrentio/tv"
	"sync"
)

func findByEpisode(details overseerr_structs.TvDetails, season int, episode int, episodeWg *sync.WaitGroup) {
	tvId := details.MediaInfo.TmdbID

	streams, err := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, episode)
	if err != nil {
		fmt.Printf("[%v - S%vE%v] Failed to get results\n", tvId, season, episode)
		fmt.Println(err)
		episodeWg.Done()
		return
	}

	streams = torrentio_tv.FilterEpisodes(streams)
	streams = torrentio.FilterVersions(streams, "shows")

	if len(streams) == 0 {
		fmt.Printf("[%v - S%vE%v] No complete episodes found\n", tvId, season, episode)
		episodeWg.Done()
		return
	}

	for _, stream := range streams {
		properties, err := torrentio.GetPropertiesFromStream(stream)
		if err != nil {
			fmt.Printf("[%v - S%vE%v - %s] Failed to get properties\n", tvId, season, episode, stream.Version)
			fmt.Println(err)
			continue
		}

		if !torrentio.MatchesProperties(stream, properties) {
			continue
		}

		fmt.Printf("[%v - S%vE%v - %s] + %v\n", tvId, season, episode, stream.Version, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Printf("[%v - S%vE%v - %s] Failed to add magnet\n", tvId, season, episode, stream.Version)
			fmt.Println(err)
			continue
		}
	}

	episodeWg.Done()
}

func findBySeason(details overseerr_structs.TvDetails, season int, seasonWg *sync.WaitGroup) {
	tvId := details.MediaInfo.TmdbID

	streams, err := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, 1)
	if err != nil {
		fmt.Printf("[%v - S%v] Failed to get results\n", tvId, season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}

	streams, err = torrentio_tv.FilterSeasons(streams)
	if err != nil {
		fmt.Printf("[%v - S%v] Failed to filter seasons\n", tvId, season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}

	streams = torrentio.FilterVersions(streams, "shows")

	if len(streams) == 0 {
		fmt.Printf("[%v - S%v] No complete seasons found, searching for episodes\n", tvId, season)

		seaonDetails, err := getSeasonDetails(details.MediaInfo.TmdbID, season)
		if err != nil {
			fmt.Printf("[%v - S%v] Failed to get details\n", tvId, season)
			fmt.Println(err)
			seasonWg.Done()
			return
		}

		episodes := getEpisodeNumbersBySeason(seaonDetails, season)

		if len(episodes) == 0 {
			fmt.Printf("[%v - S%v] Failed to get episode count\n", tvId, season)
			seasonWg.Done()
			return
		}

		var episodesWg sync.WaitGroup
		for _, episode := range episodes {
			episodesWg.Add(1)
			go findByEpisode(details, season, episode, &episodesWg)
		}

		episodesWg.Wait()
		seasonWg.Done()

		return
	}

	for _, stream := range streams {
		properties, err := torrentio.GetPropertiesFromStream(stream)
		if err != nil {
			fmt.Printf("[%v - S%v - %s] Failed to get properties\n", tvId, season, stream.Version)
			fmt.Println(err)
			continue
		}

		if !torrentio.MatchesProperties(stream, properties) {
			continue
		}

		fmt.Printf("[%v - S%v - %s] + %v\n", tvId, season, stream.Version, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Printf("[%v S%v - %s] Failed to add magnet\n", tvId, season, stream.Version)
			fmt.Println(err)
			continue
		}
	}

	seasonWg.Done()
}

func findBySeasonPlex(details overseerr_structs.TvDetails, season int, seasonWg *sync.WaitGroup) {
	tvId := details.MediaInfo.TmdbID

	showLeaves, err := plex.GetShowLeaves(details.MediaInfo.RatingKey)
	if err != nil {
		fmt.Printf("[%v - S%v] Failed to get show leaves\n", tvId, season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}

	storedEpisodes := plex.GetEpisodesBySeason(showLeaves.Video, season)

	seasonDetails, err := getSeasonDetails(details.MediaInfo.TmdbID, season)
	if err != nil {
		fmt.Printf("[%v - S%v] Failed to get season details\n", tvId, season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}

	seasonEpisodes := getEpisodeNumbersBySeason(seasonDetails, season)

	if len(seasonEpisodes) == 0 {
		fmt.Printf("[%v - S%v] Season has no episodes\n", tvId, season)
		seasonWg.Done()
		return
	}

	releasedEpisodes := getReleasedEpisodeNumbersBySeason(seasonDetails, season)

	if len(releasedEpisodes) == 0 {
		fmt.Printf("[%v - S%v] Season has no released episodes\n", tvId, season)
		seasonWg.Done()
		return
	}

	if len(storedEpisodes) == 0 {
		fmt.Printf("[%v - S%v] Season has %v released episodes\n", tvId, season, len(releasedEpisodes))

		if len(releasedEpisodes) == len(seasonEpisodes) {
			fmt.Printf("[%v - S%v] Season is fully released, searching for complete season\n", tvId, season)

			go findBySeason(details, season, seasonWg)
		} else {
			fmt.Printf("[%v - S%v] Season is not fully released, searching for episodes\n", tvId, season)

			var episodesWg sync.WaitGroup
			for _, episode := range releasedEpisodes {
				episodesWg.Add(1)
				go findByEpisode(details, season, episode, &episodesWg)
			}

			episodesWg.Wait()
		}
	} else {
		fmt.Printf("[%v - S%v] Episodes for season found on plex, searching for episodes\n", tvId, season)

		var episodesWg sync.WaitGroup
		for _, episode := range releasedEpisodes {
			for _, storedEpisode := range storedEpisodes {
				if storedEpisode.Index == fmt.Sprintf("%v", episode) {
					continue
				}
			}

			episodesWg.Add(1)
			go findByEpisode(details, season, episode, &episodesWg)
		}

		episodesWg.Wait()
	}

	seasonWg.Done()
}

func findById(tvId int, seasons []int) {
	if len(seasons) == 0 {
		fmt.Println("No seasons requested")
		return
	}

	fmt.Printf("[%v] Searching for seasons %v\n", tvId, seasons)

	details, err := getTvDetails(tvId)
	if err != nil {
		fmt.Println("Failed to get details")
		fmt.Println(err)
		return
	}

	fmt.Printf("[%v] %s\n", tvId, details.OriginalName)

	var seasonWg sync.WaitGroup

	if details.MediaInfo.RatingKey != "" {
		fmt.Printf("[%v] Content of this show is on plex so we should only download missing seasons or episodes\n", tvId)

		for _, season := range seasons {
			seasonWg.Add(1)
			go findBySeasonPlex(details, season, &seasonWg)
		}
	} else {
		fmt.Printf("[%v] No content of this show is on plex so we should download all requested seasons\n", tvId)

		for _, season := range seasons {
			seasonWg.Add(1)
			go findBySeason(details, season, &seasonWg)
		}
	}

	seasonWg.Wait()

	fmt.Printf("[%v] Finished\n", tvId)
}

func getTvDetails(tvId int) (overseerr_structs.TvDetails, error) {
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

func getSeasonDetails(tvId int, seasonId int) (overseerr_structs.Season, error) {
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
