# pod-link
Narrowed down alternative to plex debrid. Specifically combining the power of Overseerr, Torrentio and Real debrid.
## Build
```sh
go build main.go
```
Docker:
```sh
docker build -t pod-link:latest .
```

## Run
```sh
pod-link
```
Docker:
```
docker run --name pod-link --network host pod-link
```

## Configuration
### pod-link
configure the port that `pod-link` lives on and configure a header authorization code that must match the one in overseerr's webhook settings. Configure the minimum request age for missing content scanning, defauls are suggested

```yml
settings:
  pod:
    port: 42069
    authorization: "Overseerr notification webhook Authorization Header"
    missing_content:
      request_age: 24
```

### Real Debrid
fill in your [RD token](https://real-debrid.com/apitoken). Configure a timeout for adding magnets, defaults are suggested
```yml
settings:
  real_debrid:
    token: "TOKEN"
    timeout: 5
```

### Overseerr
To make `pod-link` actually do something you must set the notification webhook in overseerr to the url of the project (by default `localhost:8080/webhook`)
```yml
settings:
  overseerr:
    host: "http://localhost:5055"
    token: "TOKEN:"
```

### Plex
Get your plex library id's [here](https://plex.tv/devices.xml)
```yml
settings:
  plex:
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
    - "(?i)[. -]s\\d+[. -]"
    - "(?i)[. -]season[. -]\\d+[. -]"
  episodes:
    - "(?i)[. -]e\\d+[. -]"
    - "(?i)[. -]episode[. -]\\d+[. -]"
    - "(?i)[. ]s\\d+e\\d+[. ]"
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
Open the `config.example.yml` in the repo files

## Missing content scanning
`pod-link` can look at overseerr requests that have not been (fully) completed yet and try to complete them. This is useful for series that are still releasing new episodes, it will periodically check if a new episode has been released and add it. Another usecase would be media that has become unavailable, then it will search for that media again and mostlikely find something for in its place.

### How to use
You can trigger a scan by running `pod-link missing-content` and it will automatically do its thing and close once the job is done. Most people will probably want to run this command periodically, I recommend running a `crontab` or something similar on the machine you're running `pod-link` on.

Add to your contab using:
```sh
crontab -e
```

Example:
```sh
0 */12 * * * pod-link missing-content
```
Docker:
```sh
0 */12 * * * docker exec pod-link /app/pod-link missing-content
```


## Credits
[Plex Debrid](https://github.com/itsToggle/plex_debrid/) A lot of the inspiration

[Torrentio](https://github.com/TheBeastLT/torrentio-scraper) The source for all media

[Overseerr](https://github.com/sct/overseerr)
