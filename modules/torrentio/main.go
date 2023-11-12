package torrentio

import (
	"fmt"
	"pod-link/modules/config"
	"regexp"
	"strings"
)

type Stream struct {
	Name          string `json:"name"`
	Title         string `json:"title"`
	Url           string `json:"url"`
	Version       string
	BehaviorHints struct {
		BingeGroup string `json:"bingeGroup"`
	}
}

type Response struct {
	Streams []Stream
}

type Properties struct {
	Title   string
	Size    string
	Link    string
	Seeds   string
	Source  string
	Release string
	Files   string
}

func GetBaseURL(mediaType string) string {
	settings := config.GetSettings()

	var filter string
	switch mediaType {
	case "shows":
		filter = settings.Torrentio.Shows.FilterURI
	case "movies":
		filter = settings.Torrentio.Movies.FilterURI
	default:
		filter = ""
	}

	token := settings.RealDebrid.Token
	url := fmt.Sprintf("https://torrentio.strem.fun/%s|realdebrid=%s", filter, token)

	return url
}

func getEmojiString(input []string) string {
	for _, title := range input {
		if strings.Contains(title, "👤") && strings.Contains(title, "💾") && strings.Contains(title, "⚙️") {
			return title
		}
	}

	return ""
}

func getEmojiValues(input string) ([]string, error) {
	pattern := `👤\s(.*?)\s💾\s(.*?)\s⚙️\s(.*$)`

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Error compiling regular expression: %v\n", err)
		return nil, err
	}

	return regex.FindStringSubmatch(input)[1:], nil
}

func getMagnet(input string) (string, string) {
	input = strings.ReplaceAll(input, "https://torrentio.strem.fun/realdebrid/", "")
	split := strings.Split(input, "/")

	hash := split[1]
	so := split[3]

	if so == "undefined" || so == "null" || so == "0" {
		so = "all"
	}

	return "magnet:?xt=urn:btih:" + hash, so
}

func GetPropertiesFromStream(stream Stream) (Properties, error) {
	var properties Properties

	titleSplit := strings.Split(stream.Title, "\n")
	emojiString := getEmojiString(titleSplit)

	emojiValues, err := getEmojiValues(emojiString)
	if err != nil {
		fmt.Printf("Error getting emoji values: %v\n", err)
		return Properties{}, err
	}

	magnet, files := getMagnet(stream.Url)

	properties.Title = strings.ReplaceAll(titleSplit[0], " ", ".")
	properties.Link = magnet
	properties.Seeds = emojiValues[0]
	properties.Size = emojiValues[1]
	properties.Source = emojiValues[2]
	properties.Release = "[torrentio: " + properties.Source + "]"
	properties.Files = files

	return properties, nil
}

func getByVersion(version config.Version, streams []Stream) (Stream, error) {
	for _, stream := range streams {
		match := true

		for _, include := range version.Include {
			regex, err := regexp.Compile(include)
			if err != nil {
				fmt.Printf("Error compiling include regex for version: %v\n", version.Name)
				return Stream{}, err
			}

			if !regex.MatchString(stream.Title) {
				match = false
				break
			}
		}

		for _, exclude := range version.Exclude {
			regex, err := regexp.Compile(exclude)
			if err != nil {
				fmt.Printf("Error compiling exclude regex for version: %v\n", version.Name)
				return Stream{}, err
			}

			if regex.MatchString(stream.Title) {
				match = false
				break
			}
		}

		if match {
			return stream, nil
		}
	}

	return Stream{}, nil
}

func FilterVersions(streams []Stream, mediaType string) []Stream {
	var results []Stream

	versions := config.GetVersions(mediaType)
	for _, version := range versions {
		result, err := getByVersion(version, streams)
		if err != nil {
			fmt.Printf("Error getting version: %v\n", version.Name)
			fmt.Println(err)
			continue
		}

		if result == (Stream{}) {
			continue
		}

		result.Version = version.Name
		results = append(results, result)
	}

	if len(results) == 0 && len(streams) > 0 {
		streams[0].Version = "Fallback"
		results = append(results, streams[0])
	}

	return results
}

func MatchesProperties(stream Stream, properties Properties) bool {
	if properties.Files == "all" || properties.Files == "" {
		return true
	}

	fileCount := len(strings.Split(properties.Files, ","))
	maxFiles := config.GetConfig().Movies.MaxFiles

	if fileCount <= maxFiles {
		return true
	}

	return false
}
