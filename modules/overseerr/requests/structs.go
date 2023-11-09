package overseerr_requests

import (
	overseerr_structs "pod-link/modules/overseerr/structs"
)

type RequestsReturned struct {
	PageInfo overseerr_structs.PageInfo       `json:"pageInfo"`
	Results  []overseerr_structs.MediaRequest `json:"results"`
}
