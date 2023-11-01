package plex

import (
	"fmt"
	"net/http"
	"pod-link/modules/config"
)

func RefreshLibrary(id string) error {
	settings := config.GetSettings()
	host := settings.Plex.Host
	token := settings.Plex.Token

	url := fmt.Sprintf("%s/library/sections/%s/refresh?X-Plex-Token=%s", host, id, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create request")
		return err
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to send request")
		return err
	}

	defer response.Body.Close()

	return nil
}
