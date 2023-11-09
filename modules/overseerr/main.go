package overseerr

import (
	overseerr_settings "pod-link/modules/overseerr/settings"
	overseerr_structs "pod-link/modules/overseerr/structs"
)

func getServerConnection(connections []overseerr_structs.PlexConnection) (overseerr_structs.PlexConnection, error) {
	for _, connection := range connections {
		if connection.Status == 200 {
			return connection, nil
		}
	}

	return overseerr_structs.PlexConnection{}, nil
}

func GetPlexTokenAndHost() (string, string, error) {
	plexSettings, err := overseerr_settings.GetPlexSettings()
	if err != nil {
		return "", "", err
	}

	machineId := plexSettings.MachineID

	plexServers, err := overseerr_settings.GetPlexServers()
	if err != nil {
		return "", "", err
	}

	for _, server := range plexServers {
		if server.ClientIdentifier == machineId {
			connection, err := getServerConnection(server.Connection)
			if err != nil {
				return "", "", err
			}

			return server.AccessToken, connection.Uri, nil
		}
	}

	return "", "", nil
}
