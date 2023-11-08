package overseerr_tv

import (
	"fmt"
	"pod-link/modules/debrid"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"pod-link/modules/torrentio"
	torrentio_tv "pod-link/modules/torrentio/tv"
	"strconv"
	"sync"
)

func FindByEpisode(season int, episode int, details overseerr_structs.TvDetails, wg *sync.WaitGroup) {
	streams, err := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, episode)
	if err != nil {
		fmt.Printf("[S%vE%v] Failed to get results\n", season, episode)
		fmt.Println(err)
		wg.Done()
		return
	}

	streams = torrentio_tv.FilterEpisodes(streams)
	streams = torrentio.FilterVersions(streams, "shows")

	if len(streams) == 0 {
		fmt.Printf("[S%vE%v] Not found\n", season, episode)
		wg.Done()
		return
	}

	for _, stream := range streams {
		properties, err := torrentio.GetPropertiesFromStream(stream)
		if err != nil {
			fmt.Printf("[%s - S%vE%v] Failed to get properties\n", stream.Version, season, episode)
			fmt.Println(err)
			continue
		}

		if !torrentio.MatchesProperties(stream, properties) {
			continue
		}

		fmt.Printf("[%s - S%vE%v] + %v\n", stream.Version, season, episode, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Printf("[%s - S%vE%v] Failed to add magnet\n", stream.Version, season, episode)
			fmt.Println(err)
			continue
		}
	}

	wg.Done()
}

func FindBySeason(season int, details overseerr_structs.TvDetails, seasonWg *sync.WaitGroup) {
	streams, err := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, 1)
	if err != nil {
		fmt.Printf("[S%v] Failed to get results\n", season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}


	streams, err = torrentio_tv.FilterSeasons(streams)
	if err != nil {
		fmt.Printf("[S%v] Failed to filter seasons\n", season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}

	streams = torrentio.FilterVersions(streams, "shows")

	if len(streams) == 0 {
		fmt.Printf("[S%v] No complete seasons found, searching for episodes\n", season)

		episodes, err := getEpisodeCountBySeason(details.ID, season)
		if err != nil {
			fmt.Printf("[S%v] Failed to get episode count\n", season)
			fmt.Println(err)
			seasonWg.Done()
			return
		}

		if len(episodes) == 0 {
			fmt.Printf("[S%v] Failed to get episode count\n", season)
			seasonWg.Done()
			return
		}

		var episodesWg sync.WaitGroup
		for _, episode := range episodes {
			episodesWg.Add(1)
			go FindByEpisode(season, episode, details, &episodesWg)
		}

		episodesWg.Wait()
		seasonWg.Done()
		return
	}

	for _, stream := range streams {
		properties, err := torrentio.GetPropertiesFromStream(stream)
		if err != nil {
			fmt.Printf("[%s - S%v] Failed to get properties\n", stream.Version, season)
			fmt.Println(err)
			continue
		}

		if !torrentio.MatchesProperties(stream, properties) {
			continue
		}

		fmt.Printf("[%s - S%v] + %v\n", stream.Version, season, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Printf("[%s - S%v] Failed to add magnet\n", stream.Version, season)
			fmt.Println(err)
			continue
		}
	}

	seasonWg.Done()
}

func FindById(tvId int, seasons []int) {
	details, err := GetTvDetails(tvId)
	if err != nil {
		fmt.Println("Failed to get details")
		fmt.Println(err)
		return
	}

	fmt.Printf("[%v] %s\n", details.MediaInfo.TmdbID, details.OriginalName)

	var seasonWg sync.WaitGroup
	for _, season := range seasons {
		seasonWg.Add(1)
		go FindBySeason(season, details, &seasonWg)
	}

	seasonWg.Wait()
}


func Request(notification overseerr_structs.MediaAutoApprovedNotification) {
	TmdbId, err := strconv.Atoi(notification.Media.TmdbId)
	if err != nil {
		fmt.Println("Failed to convert tmdb id to int")
		fmt.Println(err)
		return
	}

	seasons := getRequestedSeasons(notification.Extra)

	if len(seasons) == 0 {
		fmt.Println("No seasons requested")
		return
	}

	fmt.Println("Requested seasons:", seasons)

	FindById(TmdbId, seasons)
}
