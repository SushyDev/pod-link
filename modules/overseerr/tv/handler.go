package tv

import (
	"fmt"
	"pod-link/modules/config"
	"pod-link/modules/debrid"
	"pod-link/modules/plex"
	"pod-link/modules/structs"
	"pod-link/modules/torrentio"
	torrentio_tv "pod-link/modules/torrentio/tv"
	"sync"
	"time"
)

func FindByEpisode(season int, episode int, details Tv, wg *sync.WaitGroup) {
	results, err := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, episode)
	if err != nil {
		fmt.Printf("[S%vE%v] Failed to get results\n", season, episode)
		fmt.Println(err)
		wg.Done()
		return
	}

	episodes := torrentio_tv.FilterEpisodes(results)
	filtered := torrentio.FilterVersions(episodes, "shows")

	if len(filtered) == 0 {
		fmt.Printf("[S%vE%v] Not found\n", season, episode)
		wg.Done()
		return
	}

	for _, result := range filtered {
		properties, err := torrentio.GetPropertiesFromStream(result)
		if err != nil {
			fmt.Printf("[%s - S%vE%v] Failed to get properties\n", result.Version, season, episode)
			fmt.Println(err)
			continue
		}

		fmt.Printf("[%s - S%vE%v] + %v\n", result.Version, season, episode, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Printf("[%s - S%vE%v] Failed to add magnet\n", result.Version, season, episode)
			fmt.Println(err)
			continue
		}
	}

	wg.Done()
}

func FindBySeason(season int, details Tv, seasonWg *sync.WaitGroup) {
	results, err := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, 1)
	if err != nil {
		fmt.Printf("[S%v] Failed to get results\n", season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}


	seasons, err := torrentio_tv.FilterSeasons(results)
	if err != nil {
		fmt.Printf("[S%v] Failed to filter seasons\n", season)
		fmt.Println(err)
		seasonWg.Done()
		return
	}

	filtered := torrentio.FilterVersions(seasons, "shows")

	if len(filtered) == 0 {
		fmt.Printf("[S%v] No complete seasons found, searching for episodes\n", season)
		episodes := getEpisodeCountBySeason(season, details.Seasons)

		if episodes == 0 {
			fmt.Printf("[S%v] Failed to get episode count\n", season)
			seasonWg.Done()
			return
		}

		var episodesWg sync.WaitGroup
		for episode := 1; episode <= episodes; episode++ {
			episodesWg.Add(1)
			go FindByEpisode(season, episode, details, &episodesWg)
		}

		episodesWg.Wait()
		seasonWg.Done()
		return
	}

	for _, result := range filtered {
		properties, err := torrentio.GetPropertiesFromStream(result)
		if err != nil {
			fmt.Printf("[%s - S%v] Failed to get properties\n", result.Version, season)
			fmt.Println(err)
			continue
		}

		fmt.Printf("[%s - S%v] + %v\n", result.Version, season, properties.Title)

		err = debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Printf("[%s - S%v] Failed to add magnet\n", result.Version, season)
			fmt.Println(err)
			continue
		}
	}

	seasonWg.Done()
}

func Request(notification structs.MediaAutoApprovedNotification) {
	details, err := GetDetails(notification.Media.TmdbId)
	if err != nil {
		fmt.Println("Failed to get details")
		fmt.Println(err)
		return
	}

	fmt.Println("Got request for", details.Name)

	seasons := getRequestedSeasons(notification.Extra)
	fmt.Println("Requested seasons:", seasons)

	if len(seasons) == 0 {
		fmt.Println("No seasons found")
		return
	}

	var seasonWg sync.WaitGroup
	for _, season := range seasons {
		seasonWg.Add(1)
		go FindBySeason(season, details, &seasonWg)
	}

	seasonWg.Wait()

	settings := config.GetSettings()
	host := settings.Plex.Host
	token := settings.Plex.Token
	tvId := settings.Plex.TvId

	if host != "" && token != "" && tvId != "" {
		time.Sleep(20 * time.Second)
		err := plex.RefreshLibrary(tvId)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to refresh library")
		}
	}
}
