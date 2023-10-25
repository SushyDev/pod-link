package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func validateEnv() {
	if os.Getenv("REAL_DEBRID_TOKEN") == "" {
		panic("REAL_DEBRID_TOKEN is not set")
	}

	if os.Getenv("OVERSEERR_HOST") == "" {
		panic("OVERSEERR_HOST is not set")
	}

	if os.Getenv("OVERSEERR_TOKEN") == "" {
		panic("OVERSEERR_TOKEN is not set")
	}

	if (os.Getenv("PLEX_HOST") == "" || os.Getenv("PLEX_TOKEN") == "") && (os.Getenv("PLEX_HOST") != "" || os.Getenv("PLEX_TOKEN") != "") {
		panic("PLEX_HOST and PLEX_TOKEN must both be set or both be empty")
	}

	if os.Getenv("PLEX_HOST") != "" && os.Getenv("PLEX_TOKEN") != "" {
		if os.Getenv("PLEX_TV_ID") == "" {
			panic("PLEX_TV_ID is not set")
		}

		if os.Getenv("PLEX_MOVIE_ID") == "" {
			panic("PLEX_MOVIE_ID is not set")
		}
	}
}

func main() {
	validateEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		var requestData RequestData
		err = json.Unmarshal(body, &requestData)
		if err != nil {
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
