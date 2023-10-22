package tv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project-pod/modules/torrentio"
)

func GetList(ImdbId string, Season int, Episode int) []torrentio.Stream {
	filter := "sort=qualitysize%7Cqualityfilter=other,scr,cam,unknown"
	url := fmt.Sprintf("https://torrentio.strem.fun/%s/stream/series/%s:%v:%v.json", filter, ImdbId, Season, Episode)

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
