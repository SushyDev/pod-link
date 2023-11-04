package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pod-link/modules/config"
	overseerr "pod-link/modules/overseerr"
	"pod-link/modules/structs"
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
        var mediaAutoApprovedNotification structs.MediaAutoApprovedNotification
        err := json.Unmarshal(body, &mediaAutoApprovedNotification)
        if err != nil {
            return err
        }

        overseerr.HandleMediaAutoApprovedNotification(mediaAutoApprovedNotification)
    default:
        fmt.Printf("Unknown notification type: %s\n", notificationType)
    }

    return nil
}
