package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pod-link/modules/config"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_settings "pod-link/modules/overseerr/settings"
	overseerr_structs "pod-link/modules/overseerr/structs"
	overseerr_tv "pod-link/modules/overseerr/tv"
	"pod-link/modules/plex"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pod-link up and running!"))
		r.Body.Close()
	})

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
		var notification overseerr_structs.MediaAutoApprovedNotification
		err := json.Unmarshal(body, &notification)
		if err != nil {
			return err
		}

		switch notification.Media.MediaType {
		case "movie":
			overseerr_movies.Request(notification)
		case "tv":
			overseerr_tv.Request(notification)
		}

		libraryIds := overseerr_settings.GetLibraryIdsByType(notification.Media.MediaType)
		for _, libraryId := range libraryIds {
			plex.RefreshLibrary(libraryId)
		}
	default:
		fmt.Printf("Unknown notification type: %s\n", notificationType)
	}

	return nil
}
