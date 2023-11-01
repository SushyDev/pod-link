package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pod-link/modules/config"
	overseerr_movies "pod-link/modules/overseerr/movies"
	overseerr_tv "pod-link/modules/overseerr/tv"
	"pod-link/modules/structs"
)

func handleMediaAutoApprovedNotification(notification structs.MediaAutoApprovedNotification) {
	switch notification.Media.MediaType {
	case "movie":
		overseerr_movies.Request(notification)
	case "tv":
		overseerr_tv.Request(notification)
	}
}

type RequestData struct {
	NotificationType string `json:"notification_type"`
}

func main() {
	settings := config.GetSettings()
	port := settings.Pod.Port
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println("Only POST requests are allowed")
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}
	
		if r.Header.Get("Authorization") != settings.Pod.Authorization {
			fmt.Println("Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		var requestData RequestData
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
			return
		}

		switch requestData.NotificationType {
		// case "MEDIA_APPROVED":
		//     var mediaApprovedNotification structs.MediaApprovedNotification
		//     err = json.Unmarshal(body, &mediaApprovedNotification)
		//     if err != nil {
		//         http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		//         return
		//     }
		//
		//     handleMediaAutoApprovedNotification(mediaApprovedNotification.MediaAutoApprovedNotification)
		case "MEDIA_AUTO_APPROVED":
			var mediaAutoApprovedNotification structs.MediaAutoApprovedNotification
			err = json.Unmarshal(body, &mediaAutoApprovedNotification)
			if err != nil {
				http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
				return
			}

			handleMediaAutoApprovedNotification(mediaAutoApprovedNotification)
		default:
			fmt.Println("unknown")
		}

		responseData, err := json.Marshal(requestData)
		if err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err = w.Write(responseData)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

		fmt.Println("Finished!")
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
