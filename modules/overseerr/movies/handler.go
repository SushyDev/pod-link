package overseerr_movies

import (
	"fmt"
	overseerr_structs "pod-link/modules/overseerr/structs"
	"strconv"
)

func Missing(requestDetails overseerr_structs.MediaRequest) {
	FindById(requestDetails.Media.TmdbID)
}

func Request(notification overseerr_structs.MediaAutoApprovedNotification) {
	movieId, err := strconv.Atoi(notification.Media.TmdbId)
	if err != nil {
		fmt.Println("Failed to convert tmdb id to int")
		fmt.Println(err)
		return
	}

	FindById(movieId)
}
