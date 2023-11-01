package anime

import (
	"fmt"
	"pod-link/modules/structs"
)

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
}
