package handler

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    overseerr_movies "project-pod/modules/overseerr/movies"
    overseerr_tv "project-pod/modules/overseerr/tv"
    "project-pod/modules/structs"
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

func Handler(w http.ResponseWriter, r *http.Request) {
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

    // Marshal the modified JSON back to bytes
    responseData, err := json.Marshal(requestData)
    if err != nil {
        http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
        return
    }

    // Set response headers
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    fmt.Println("\nFinished")

    // Write the modified JSON to the response
    _, err = w.Write(responseData)
    if err != nil {
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
    }
}
