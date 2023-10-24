# pod-link
Narrowed down alternative to plex debrid. Specifically combining the power of Overseerr, Torrentio and Real debrid.
## Build
`go build main.go`

## Run
Run the executable and make sure all these environment variables are available:
```
REAL_DEBRID_API_KEY=
OVERSEERR_HOST=
OVERSEERR_TOKEN=
PLEX_HOST=
PLEX_TOKEN=
PLEX_TV_ID=
PLEX_MOVIE_ID=
```
PLEX_X_ID= is for the id of the corresponding library in plex

# Configuration
To make `pod-link` actually do something you must set the notification webhook in overseerr to the url of the project (by default `localhost:8080/webhook`)

## Credits
[Plex Debrid](https://github.com/itsToggle/plex_debrid/) A lot of the inspiration

[Torrentio](https://github.com/TheBeastLT/torrentio-scraper) The source for all media

[Overseerr](https://github.com/sct/overseerr)
