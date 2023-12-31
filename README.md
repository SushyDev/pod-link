# pod-link
Narrowed down alternative to plex debrid. Specifically combining the power of Overseerr, Torrentio and Real debrid.
## Build
`go build main.go`

# Configuration
### pod-link
you can configure `pod-link`'s settings:
```yml
settings:
  pod:
    port: 42069
    authorization: "Overseerr notification webhook Authorization Header"
```

### Real Debrid
fill in your [RD token](https://real-debrid.com/apitoken):
```yml
settings:
  real_debrid:
    token: "TOKEN"
```

### Overseerr
```yml
settings:
  overseerr:
    host: "http://localhost:5055"
    token: "TOKEN:"
```

### Plex
Get your [plex token](https://github.com/SushyDev/plex-oauth) or [here](https://plex.tv/devices.xml) and get your libary id's
```yml
settings:
  plex:
    host: "http://localhost:32400"
    token: "TOKEN"
    tv_id: 1
    movie_id: 2
```

### Torrentio
Get the torrentio filter options from [here](https://torrentio.strem.fun/configure)
Make sure to only put the filter options and not the entire url.
**Don't** configure Real Debrid (or any other debrid) in Torrentio, `pod-link` will do that for you!
```yml
settings:
  torrentio:
    shows:
      filter_uri: "sort=qualitysize|qualityfilter=other,scr,cam,unknown"
    movies:
      filter_uri: "sort=qualitysize|qualityfilter=other,scr,cam,unknown"
```

### Shows
Here you can configure the regex to determine what `pod-link` should consider a complete season or just a single episode
Defaults supplied in the example should suffice for most usecases
```yml
shows:
  seasons:
    - "(?i)[. ]s\\d+[. ]"
    - "(?i)[. ]season \\d+[. ]"
  episodes:
    - "(?i)[. ]e\\d+[. ]"
    - "(?i)[. ]episode \\d+[. ]"
```

### Movies
Here you can configure details about what results should be used for movies. For now the only thing you can configure is the max file count that a movie can have, you can use this to prevent `pod-link` from picking up results which ship lots of m2ts containers
```yml
movies:
  max_files: 10
```

### Versions
Here you can configure all the versions of a movie/season/episode to download.
You can configure them per media type or for all media types.
Media types:
- all
- movies
- shows
- anime (not implemented yet)

As you can imagine all the versions in "all" will apply to all media types.
Within each media type you can configure the version, each version must have a unique name.
If you name a version "all" its include and exclude will be appended to all other versions for that media type.
A version must have a name and can have either or both a list of include and exclude regex strings.
Regex is handled by golang's default regex implementation so any limitations there will apply here.

### Example config
Open the config.yml in the repo files

To make `pod-link` actually do something you must set the notification webhook in overseerr to the url of the project (by default `localhost:8080/webhook`)

## Credits
[Plex Debrid](https://github.com/itsToggle/plex_debrid/) A lot of the inspiration

[Torrentio](https://github.com/TheBeastLT/torrentio-scraper) The source for all media

[Overseerr](https://github.com/sct/overseerr)
