package config

type Pod struct {
    Port string `yaml:"port"`
    Authorization string `yaml:"authorization"`
}

type RealDebrid struct {
    Token string `yaml:"token"`
}

type Overseerr struct {
    Host string `yaml:"host"`
    Token string `yaml:"token"`
}

type Plex struct {
    Host string `yaml:"host"`
    Token string `yaml:"token"`
    TvId string `yaml:"tv_id"`
    MovieId string `yaml:"movie_id"`
}

type Torrentio struct {
    FilterURI string `yaml:"filter_uri"`
}

type Settings struct {
    Pod Pod `yaml:"pod"`
    RealDebrid RealDebrid `yaml:"real_debrid"`
    Overseerr Overseerr `yaml:"overseerr"`
    Plex Plex `yaml:"plex"`
    Torrentio Torrentio `yaml:"torrentio"`
}

type Version struct {
    Name string `yaml:"name"`
    Include []string `yaml:"include"`
    Exclude []string `yaml:"exclude"`
}

type Versions struct {
    All []Version `yaml:"all"`
    Movie []Version `yaml:"movie"`
    Show []Version `yaml:"show"`
}

type Shows struct {
    Seasons []string `yaml:"seasons"`
    Episodes []string `yaml:"episodes"`
}

type Config struct {
    Settings Settings `yaml:"settings"`
    Shows Shows `yaml:"shows"`
    Versions Versions `yaml:"versions"`
}

