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

func GetBaseURL() string {
	settings := config.GetSettings()
	filter := settings.Torrentio.FilterURI
	token := settings.RealDebrid.Token
	url := fmt.Sprintf("https://torrentio.strem.fun/%s|realdebrid=%s", filter, token)

	return url
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
	split := strings.Split(input, "/")
	hash := split[5]
	so := split[6]
	dn := split[8]

	if so == "null" {
		so = "all"
	}

	return "magnet:?xt=urn:btih:" + hash + "&dn=" + dn, so
}

func GetPropertiesFromStream(stream Stream) Properties {
	var properties Properties

	emojiString := ""

	titleSplit := strings.Split(stream.Title, "\n")
	for _, title := range titleSplit {
		if strings.Contains(title, "👤") && strings.Contains(title, "💾") && strings.Contains(title, "⚙️") {
			emojiString = title
		}
	}

	emojiValues, err := getEmojiValues(emojiString)
	if err != nil {
		fmt.Printf("Error getting emoji values: %v\n", err)
		return properties
	}

	magnet, files := getMagnet(stream.Url)

	properties.Title = strings.ReplaceAll(titleSplit[0], " ", ".")
	properties.Link = magnet
	properties.Seeds = emojiValues[0]
	properties.Size = emojiValues[1]
	properties.Source = emojiValues[2]
	properties.Release = "[torrentio: " + properties.Source + "]"
	properties.Files = files

	return properties
}

func getByVersion(version config.Version, streams []Stream) Stream {
	for _, stream := range streams {
		match := true

		for _, include := range version.Include {
			matched, err := regexp.MatchString(include, stream.Title)
			if err != nil {
				fmt.Printf("Error matching regular expression: %v\n", err)
				return Stream{}
			}

			if !matched {
				match = false
				break
			}
		}

		for _, exclude := range version.Exclude {
			matched, err := regexp.MatchString(exclude, stream.Title)
			if err != nil {
				fmt.Printf("Error matching regular expression: %v\n", err)
				return Stream{}
			}

			if matched {
				match = false
				break
			}
		}

		if match {
			return stream
		}
	}

	return Stream{}
}

func FilterFormats(streams []Stream, mediaType string) []Stream {
	var results []Stream

	versions := config.GetVersions(mediaType)

	for _, version := range versions {
		result := getByVersion(version, streams)
		if result != (Stream{}) {
			result.Version = version.Name
			results = append(results, result)
		}
	}

	if len(results) == 0 && len(streams) > 0 {
		streams[0].Version = "Fallback"
		results = append(results, streams[0])
	}

	return results
}
