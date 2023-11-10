package debrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"pod-link/modules/config"
)

type AddMagnetResponse struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

func AddMagnet(magnet string, files string) error {
	input := url.Values{}
	input.Set("magnet", magnet)

	requestBody := input.Encode()
	req, err := http.NewRequest("POST", "https://api.real-debrid.com/rest/1.0/torrents/addMagnet", bytes.NewBufferString(requestBody))
	if err != nil {
		fmt.Println("Failed to create request")
		return err
	}

	settings := config.GetSettings()
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", settings.RealDebrid.Token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return err
	}

	defer response.Body.Close()

	var data AddMagnetResponse
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode response")

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to read response body")
			return err
		}

		fmt.Println(body)

		return err
	}

	switch response.StatusCode {
	case 201:
		return selectFiles(data.Id, files)
	case 400:
		return fmt.Errorf("Bad Request (see error message)")
	case 401:
		return fmt.Errorf("Bad token (expired, invalid)")
	case 403:
		return fmt.Errorf("Permission denied (account locked, not premium) or Infringing torrent")
	case 503:
		return fmt.Errorf("Service unavailable (see error message)")
	default:
		return fmt.Errorf("Unknown error")
	}
}

func deleteFile(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/delete/%s", id), nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return err
	}

	settings := config.GetSettings()
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", settings.RealDebrid.Token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case 204:
		return nil
	case 401:
		return fmt.Errorf("Bad token (expired, invalid)")
	case 403:
		return fmt.Errorf("Permission denied (account locked, not premium)")
	case 404:
		return fmt.Errorf("Unknown ressource (invalid id)")
	default:
		return fmt.Errorf("Unknown error")
	}
}

func selectFiles(id string, files string) error {
	input := url.Values{}
	input.Set("files", files)

	requestBody := input.Encode()
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/selectFiles/%s", id), bytes.NewBufferString(requestBody))
	if err != nil {
		fmt.Println("Failed to create request")
		return err
	}

	settings := config.GetSettings()
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", settings.RealDebrid.Token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request")
		return err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case 202:
		return fmt.Errorf("Action already done")
	case 204:
		return nil
	case 400:
		return fmt.Errorf("Bad Request (see error message)")
	case 401:
		return fmt.Errorf("Bad token (expired, invalid)")
	case 403:
		return fmt.Errorf("Permission denied (account locked, not premium)")
	case 404:
		err := deleteFile(id)
		if err != nil {
			return err
		}

		return fmt.Errorf("Wrong parameter (invalid file id(s)) / Unknown ressource (invalid id)")
	default:
		return fmt.Errorf("Unknown error")
	}
}
