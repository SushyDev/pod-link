package overseerr_tv

import (
	"fmt"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"strconv"
)

func Missing(requestDetails overseerr_structs.MediaRequest) {
	seasons := filterCompleteSeasons(requestDetails)

	findById(requestDetails.Media.TmdbID, seasons)
}

func Request(notification overseerr_structs.MediaAutoApprovedNotification) {
	TmdbId, err := strconv.Atoi(notification.Media.TmdbId)
	if err != nil {
		fmt.Println("Failed to convert tmdb id to int")
		fmt.Println(err)
		return
	}

	seasons := getRequestedSeasons(notification.Extra)

	findById(TmdbId, seasons)
}
