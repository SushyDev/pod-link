package tv

import (
    "fmt"
    "os"
    "project-pod/modules/debrid"
    "project-pod/modules/plex"
    "project-pod/modules/structs"
    "project-pod/modules/torrentio"
    torrentio_tv "project-pod/modules/torrentio/tv"
    "strconv"
    "strings"
)

func isAnime(keywords []Keyword) bool {
    for _, keyword := range keywords {
        if strings.ToLower(keyword.Name) == "anime" {
            return true
        }
    }

    return false
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

func getEpisodeCountBySeason(number int, seasons []Season) int {
    for _, season := range seasons {
        if season.SeasonNumber == number {
            return season.EpisodeCount
        }
    }

    return 0
}


func Request(notification structs.MediaAutoApprovedNotification) {
    details := GetDetails(notification.Media.TmdbId)

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

    requestedSeasons := getRequestedSeasons(notification.Extra)
    for _, season := range requestedSeasons {
        fmt.Println("\nSeason:", season)

        episodeCount := getEpisodeCountBySeason(season, details.Seasons)
        for episode := 1; episode <= episodeCount; episode++ {
            fmt.Println("- Episode:", episode)

            results := torrentio_tv.GetList(details.ExternalIds.ImdbID, season, episode)
            for _, result := range results {
                properties := torrentio.GetPropertiesFromStream(result)
                fmt.Println("  *", properties.Title)

                err := debrid.AddMagnet(result.InfoHash)
                if err != nil {
                    fmt.Println("\033[31m", err, "\033[0m")
                }
            }
        }
    }

    err := plex.RefreshLibrary(os.Getenv("PLEX_TV_ID"))
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to refresh library")
    }
}
