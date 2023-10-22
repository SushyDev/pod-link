package debrid

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type AddMagnetResponse struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

func AddMagnet(magnet string) error {
	input := url.Values{}
	input.Set("magnet", magnet)

	requestBody := input.Encode()
	req, err := http.NewRequest("POST", "https://api.real-debrid.com/rest/1.0/torrents/addMagnet", bytes.NewBufferString(requestBody))
	if err != nil {
		fmt.Println("Failed to create request")
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("REAL_DEBRID_API_KEY")))
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
		return err
	}

	switch response.StatusCode {
	case 201:
		return selectFiles(data.Id)
	case 400:
		return errors.New("Bad Request (see error message)")
	case 401:
		return errors.New("Bad token (expired, invalid)")
	case 403:
		return errors.New("Permission denied (account locked, not premium)")
	case 503:
		return errors.New("Service unavailable (see error message)")
	default:
		return errors.New("Unknown error")
	}
}

func deleteFile(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/delete/%s", id), nil)
	if err != nil {
		fmt.Println("Failed to create request")
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("REAL_DEBRID_API_KEY")))
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
		return errors.New("Bad token (expired, invalid)")
	case 403:
		return errors.New("Permission denied (account locked, not premium)")
	case 404:
		return errors.New("Unknown ressource (invalid id)")
	default:
		return errors.New("Unknown error")
	}
}


func selectFiles(id string) error {
	input := url.Values{}
	input.Set("files", "all")

	requestBody := input.Encode()
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/selectFiles/%s", id), bytes.NewBufferString(requestBody))
	if err != nil {
		fmt.Println("Failed to create request")
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("REAL_DEBRID_API_KEY")))
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
		return errors.New("Action already done")
	case 204:
		return nil
	case 400:
		return errors.New("Bad Request (see error message)")
	case 401:
		return errors.New("Bad token (expired, invalid)")
	case 403:
		return errors.New("Permission denied (account locked, not premium)")
	case 404:
		err := deleteFile(id)
		if err != nil {
			return err
		}

		return errors.New("Wrong parameter (invalid file id(s)) / Unknown ressource (invalid id)")
	default:
		return errors.New("Unknown error")
	}
}
