package overseerr_settings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
	overseerr_structs "pod-link/modules/overseerr/structs"
)

func GetPlexSettings() (overseerr_structs.PlexSettings, error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/settings/plex", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return overseerr_structs.PlexSettings{}, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return overseerr_structs.PlexSettings{}, err
	}

	defer response.Body.Close()

	var data overseerr_structs.PlexSettings
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return overseerr_structs.PlexSettings{}, err
	}

	return data, nil
}

func GetPlexServers() (([]overseerr_structs.PlexDevice), error) {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/settings/plex/devices/servers", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return nil, err
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return nil, err
	}

	defer response.Body.Close()

	var data []overseerr_structs.PlexDevice
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return nil, err
	}

	return data, nil
}
