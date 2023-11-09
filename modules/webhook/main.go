package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pod-link/modules/config"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_structs "pod-link/modules/overseerr/structs"
	overseerr_tv "pod-link/modules/overseerr/tv"
)

type RequestData struct {
	NotificationType string `json:"notification_type"`
}

func Listen() {
	settings := config.GetSettings()
	port := settings.Pod.Port
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println("Only POST requests are allowed")
			return
		}

		if r.Header.Get("Authorization") != settings.Pod.Authorization {
			fmt.Println("Unauthorized")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var requestData RequestData
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = handleNotification(requestData.NotificationType, body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Finished!")
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleNotification(notificationType string, body []byte) error {
	switch notificationType {
	case "MEDIA_AUTO_APPROVED":
		var mediaAutoApprovedNotification overseerr_structs.MediaAutoApprovedNotification
		err := json.Unmarshal(body, &mediaAutoApprovedNotification)
		if err != nil {
			return err
		}

		switch mediaAutoApprovedNotification.Media.MediaType {
		case "movie":
			overseerr_movies.Request(mediaAutoApprovedNotification)
		case "tv":
			overseerr_tv.Request(mediaAutoApprovedNotification)
		}
	default:
		fmt.Printf("Unknown notification type: %s\n", notificationType)
	}

	return nil
}