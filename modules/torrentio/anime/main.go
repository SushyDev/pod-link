package anime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project-pod/modules/torrentio"
)

func GetList(KitsuId string, Episode int) []torrentio.Stream {
	url := fmt.Sprintf("https://torrentio.strem.fun/stream/anime/kitsu:%s:%v.json", KitsuId, Episode)

	fmt.Println(url)

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

	return torrentio.FilterResults(data.Streams)
}
