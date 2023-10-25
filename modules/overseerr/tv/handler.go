package tv

import (
	"fmt"
	"os"
	"pod-link/modules/debrid"
	"pod-link/modules/plex"
	"pod-link/modules/structs"
	"pod-link/modules/torrentio"
	torrentio_tv "pod-link/modules/torrentio/tv"
	"strconv"
	"strings"
	"sync"
)

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

func FindByEpisode(season int, episode int, details Tv, wg *sync.WaitGroup) {
    results := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, episode)
    episodes := torrentio_tv.FilterEpisodes(results)
    filtered := torrentio.FilterFormats(episodes)

    if len(filtered) == 0 {
        fmt.Println("  [Season:", season, "] [Episode:", episode, "] No complete episodes found")
        wg.Done()
        return
    }

    for _, result := range filtered {
        properties := torrentio.GetPropertiesFromStream(result)
        fmt.Println("  [Season:", season, "] [Episode:", episode, "] *", properties.Title)

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
    filtered := torrentio.FilterFormats(seasons)

    if len(filtered) == 0 {
        fmt.Println("  [Season:", season, "] No complete seasons found, searching for episodes")
        episodes := getEpisodeCountBySeason(season, details.Seasons)

        if episodes == 0 {
            fmt.Println("  [Season:", season, "] No episodes found")
            seasonWg.Done()
            return
        }

        var episodesWg sync.WaitGroup
        for episode := 1; episode <= episodes; episode++ {
            episodesWg.Add(1)
            go FindByEpisode(season, episode, details, &episodesWg)
        }

        episodesWg.Done()
        seasonWg.Done()
        return
    }

    fmt.Println("  [Season:", season, "] Found complete season")
    for _, result := range filtered {
        properties := torrentio.GetPropertiesFromStream(result)
        fmt.Println("  [Season:", season, "] *", properties.Title)

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

    // TODO: Check kitsu if format is (name-season, episode) or (name, (all episodes before + episode))
    // if isAnime(details.Keywords) {
    //     for _, season := range requestedSeasons {
    //         if season == 0 {
    //             continue
    //         }
    //
    //         kitsuName := strings.ReplaceAll(details.Name, " ", "-")
    //         kitsuName = strings.ToLower(kitsuName)
    //
    //         if season > 1 {
    //             kitsuName = fmt.Sprintf("%s-%v", kitsuName, season)
    //         }
    //
    //         kitsuDetails := kitsu.GetDetails(kitsuName)
    //
    //         for episode := 1; episode <= kitsuDetails.Attributes.EpisodeCount; episode++ {
    //             results := torrentio_anime.GetList(kitsuDetails.ID, episode)
    //
    //
    //         }
    //     }
    // }

    seasons := getRequestedSeasons(notification.Extra)
    fmt.Println("Requested seasons:", seasons)

    if len(seasons) == 0 {
        fmt.Println("No seasons found, endning")
        return
    }

	var seasonWg sync.WaitGroup
    for _, season := range seasons {
        seasonWg.Add(1)
        go FindBySeason(season, details, &seasonWg)
    }

    seasonWg.Wait()

    if os.Getenv("PLEX_HOST") != "" && os.Getenv("PLEX_TOKEN") != "" && os.Getenv("PLEX_TV_ID") != "" {
        err := plex.RefreshLibrary(os.Getenv("PLEX_TV_ID"))
        if err != nil {
            fmt.Println(err)
            fmt.Println("Failed to refresh library")
        }
    }
}
