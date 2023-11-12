package overseerr_settings

import "fmt"

func GetLibraryIdsByType(libraryType string) []string {
	plexSettings, err := GetPlexSettings()
	if err != nil {
		fmt.Println("Failed to get plex settings")
		return nil
	}

	if libraryType == "tv" {
		libraryType = "show"
	}

	var libraryIds []string
	for _, library := range plexSettings.Libraries {
		if library.Type == libraryType {
			libraryIds = append(libraryIds, library.ID)
		}
	}

	return libraryIds
}
