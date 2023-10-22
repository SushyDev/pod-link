package kitsu

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type KitsuData struct {
    ID string `json:"id"`
    Attributes struct {
        EpisodeCount int `json:"episodeCount"`
    } `json:"attributes"`
}

type KitsuResponse struct {
    Data []KitsuData `json:"data"`
}

func GetDetails(name string) KitsuData {
    url := fmt.Sprintf("https://kitsu.io/api/edge/anime?filter[text]=%s", name)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to create request")
    }

    client := &http.Client{}

    response, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to send request")
    }

    defer response.Body.Close()

    var data KitsuResponse
    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to decode response")
    }

    fmt.Println("Found:", data.Data[0].ID)

    return data.Data[0]
}
