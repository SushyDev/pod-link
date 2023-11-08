package overseerr_structs

type User struct {
	ID           int         `json:"id"`
	Email        string      `json:"email"`
	Username     interface{} `json:"username"`
	PlexToken    string      `json:"plexToken"`
	PlexUsername string      `json:"plexUsername"`
	UserType     int         `json:"userType"`
	Permissions  int         `json:"permissions"`
	Avatar       string      `json:"avatar"`
	CreatedAt    string      `json:"createdAt"`
	UpdatedAt    string      `json:"updatedAt"`
	RequestCount int         `json:"requestCount"`
}
