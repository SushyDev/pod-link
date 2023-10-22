package movies

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type OverseerrResponse struct {
    ExternalIds struct {
        ImdbId string `json:"imdbId"`
    } `json:"externalIds"`
}

func GetDetails(id string) Movie {
    host := os.Getenv("OVERSEERR_HOST")
    token := os.Getenv("OVERSEERR_TOKEN")

    url := fmt.Sprintf("%s/api/v1/movie/%s", host, id)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to create request")
    }

    req.Header.Add("X-Api-Key", token)

    client := &http.Client{}

    response, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to send request")
    }

    defer response.Body.Close()

    var details Movie
    err = json.NewDecoder(response.Body).Decode(&details)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to decode response")
    }

    return details
}
