package torrentio_anime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	"pod-link/modules/torrentio"
)

func GetList(KitsuId string, Episode int) []torrentio.Stream {
	settings := config.GetSettings()
	filter := settings.Torrentio.FilterURI
	url := fmt.Sprintf("https://torrentio.strem.fun/%s/stream/anime/kitsu:%s:%v.json", filter, KitsuId, Episode)

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

	var data torrentio.Response
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to decode response")
	}

	return torrentio.FilterFormats(data.Streams, "anime")
}
