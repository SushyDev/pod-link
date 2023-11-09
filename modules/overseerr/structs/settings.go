package overseerr_structs

import "time"

type PlexConnection struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Uri      string `json:"uri"`
	Local    bool   `json:"local"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
}

type PlexDevice struct {
	Name                   string           `json:"name"`
	Product                string           `json:"product"`
	ProductVersion         string           `json:"productVersion"`
	Platform               string           `json:"platform"`
	PlatformVersion        string           `json:"platformVersion"`
	Device                 string           `json:"device"`
	ClientIdentifier       string           `json:"clientIdentifier"`
	CreatedAt              time.Time        `json:"createdAt"`
	LastSeenAt             time.Time        `json:"lastSeenAt"`
	Provides               []string         `json:"provides"`
	Owned                  bool             `json:"owned"`
	OwnerID                string           `json:"ownerID"`
	Home                   bool             `json:"home"`
	SourceTitle            string           `json:"sourceTitle"`
	AccessToken            string           `json:"accessToken"`
	PublicAddress          string           `json:"publicAddress"`
	HttpsRequired          bool             `json:"httpsRequired"`
	Synced                 bool             `json:"synced"`
	Relay                  bool             `json:"relay"`
	DnsRebindingProtection bool             `json:"dnsRebindingProtection"`
	NatLoopbackSupported   bool             `json:"natLoopbackSupported"`
	PublicAddressMatches   bool             `json:"publicAddressMatches"`
	Presence               bool             `json:"presence"`
	Connection             []PlexConnection `json:"connection"`
}

type PlexLibrary struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

type PlexSettings struct {
	Name      string        `json:"name"`
	MachineID string        `json:"machineId"`
	IP        string        `json:"ip"`
	Port      int           `json:"port"`
	UseSSL    bool          `json:"useSsl"`
	Libraries []PlexLibrary `json:"libraries"`
	WebAppURL string        `json:"webAppUrl"`
}
