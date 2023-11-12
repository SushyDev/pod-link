package config

type Pod struct {
	Port           string `yaml:"port"`
	Authorization  string `yaml:"authorization"`
	MissingContent struct {
		RequestAge float64 `yaml:"request_age"`
	} `yaml:"missing_content"`
}

type RealDebrid struct {
	Token   string `yaml:"token"`
	Timeout int    `yaml:"timeout"`
}

type Overseerr struct {
	Host  string `yaml:"host"`
	Token string `yaml:"token"`
}

type Torrentio struct {
	Shows struct {
		FilterURI string `yaml:"filter_uri"`
	} `yaml:"shows"`
	Movies struct {
		FilterURI string `yaml:"filter_uri"`
	} `yaml:"movies"`
}

type Settings struct {
	Pod        Pod        `yaml:"pod"`
	RealDebrid RealDebrid `yaml:"real_debrid"`
	Overseerr  Overseerr  `yaml:"overseerr"`
	Torrentio  Torrentio  `yaml:"torrentio"`
}

type Version struct {
	Name    string   `yaml:"name"`
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude"`
}

type Versions struct {
	All    []Version `yaml:"all"`
	Movies []Version `yaml:"movies"`
	Shows  []Version `yaml:"shows"`
}

type Shows struct {
	Seasons  []string `yaml:"seasons"`
	Episodes []string `yaml:"episodes"`
}

type Movies struct {
	MaxFiles int `yaml:"max_files"`
}

type Config struct {
	Settings Settings `yaml:"settings"`
	Shows    Shows    `yaml:"shows"`
	Movies   Movies   `yaml:"movies"`
	Versions Versions `yaml:"versions"`
}
