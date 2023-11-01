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
	results := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, episode)
	episodes := torrentio_tv.FilterEpisodes(results)
	filtered := torrentio.FilterFormats(episodes, "show")

	if len(filtered) == 0 {
		fmt.Println(fmt.Sprintf("[S%vE%v] Not found", season, episode))
		wg.Done()
		return
	}

	for _, result := range filtered {
		properties := torrentio.GetPropertiesFromStream(result)
		fmt.Println(fmt.Sprintf("[%s - S%vE%v] + %v", result.Version, season, episode, properties.Title))

		err := debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Println("\033[31m", err, "\033[0m")
		}
	}

	wg.Done()
}

func FindBySeason(season int, details Tv, seasonWg *sync.WaitGroup) {
	results := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, 1)
	seasons := torrentio_tv.FilterSeasons(results)
	filtered := torrentio.FilterFormats(seasons, "show")

	if len(filtered) == 0 {
		fmt.Println(fmt.Sprintf("[S%v] No complete seasons found, searching for episodes", season))
		episodes := getEpisodeCountBySeason(season, details.Seasons)

		if episodes == 0 {
			fmt.Println("[Season:", season, "] No episodes found")
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
		properties := torrentio.GetPropertiesFromStream(result)
		fmt.Println(fmt.Sprintf("[%s - S%v] + %v", result.Version, season, properties.Title))

		err := debrid.AddMagnet(properties.Link, properties.Files)
		if err != nil {
			fmt.Println("\033[31m", err, "\033[0m")
		}
	}

	seasonWg.Done()
}

func Request(notification structs.MediaAutoApprovedNotification) {
	details := GetDetails(notification.Media.TmdbId)
	fmt.Println("Got request for", details.Name)

	seasons := getRequestedSeasons(notification.Extra)
	fmt.Println("Requested seasons:", seasons)

	if len(seasons) == 0 {
		fmt.Println("No seasons found, ending")
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
