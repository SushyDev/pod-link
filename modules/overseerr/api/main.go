package overseerr_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pod-link/modules/config"
)

func GetUser() {
	settings := config.GetSettings()
	host := settings.Overseerr.Host
	token := settings.Overseerr.Token

	url := fmt.Sprintf("%s/api/v1/auth/me", host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return
	}

	req.Header.Add("X-Api-Key", token)

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return
	}

	defer response.Body.Close()

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")
		return
	}

	fmt.Println(data)
}
