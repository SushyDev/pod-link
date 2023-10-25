package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pod-link/modules/torrentio"
)

func GetList(ImdbId string) []torrentio.Stream {
	realdebrid := os.Getenv("REAL_DEBRID_TOKEN")
	filter := "sort=qualitysize|qualityfilter=other,scr,cam,unknown|realdebrid=" + realdebrid
	url := fmt.Sprintf("https://torrentio.strem.fun/%s/stream/movie/%s.json", filter, ImdbId)

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

	return torrentio.FilterFormats(data.Streams)
}
