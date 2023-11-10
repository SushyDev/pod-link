package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"reflect"
)

var config Config

func GetConfig() Config {
	if !reflect.DeepEqual(config, Config{}) {
		fmt.Println("Returning cached config")
		return config
	}

	configFile, err := os.Open("config.yml")
	if err != nil {
		fmt.Printf("Error opening config file: %v\n", err)
		panic(err)
	}

	configFileBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(configFileBytes, &config)
	if err != nil {
		fmt.Printf("Error unmarshalling config file: %v\n", err)
		panic(err)
	}

	return config
}

func GetSettings() Settings {
	return GetConfig().Settings
}

func GetVersions(mediaType string) []Version {
	config := GetConfig()

	versions := config.Versions.All

	switch mediaType {
	case "movies":
		versions = append(versions, config.Versions.Movies...)
	case "shows":
		versions = append(versions, config.Versions.Shows...)
	}

	for i, iVersion := range versions {
		if iVersion.Name != "all" {
			continue
		}

		for j, jVersion := range versions {
			if jVersion.Name == "all" {
				continue
			}

			versions[j].Include = append(versions[j].Include, iVersion.Include...)
			versions[j].Exclude = append(versions[j].Exclude, iVersion.Exclude...)
		}

		if i == len(versions)-1 {
			versions = versions[:i]
			break
		}

		versions = append(versions[:i], versions[i+1:]...)
	}

	return versions
}
