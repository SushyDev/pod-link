package torrentio

import (
	"fmt"
	"regexp"
	"strings"
)

type Stream struct {
    Name          string `json:"name"`
    Title         string `json:"title"`
    Url           string `json:"url"`
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

func extractEmojiValues(input string) []string {
    pattern := `👤\s(.*?)\s💾\s(.*?)\s⚙️\s(.*$)`

    regex, err := regexp.Compile(pattern)
    if err != nil {
        fmt.Printf("Error compiling regular expression: %v\n", err)
        return nil
    }

    return regex.FindStringSubmatch(input)[1:]
}

func parseLink(input string) (string, string) {
    split := strings.Split(input, "/")
    hash := split[5]
    so := split[6]
    dn := split[8]

    return "magnet:?xt=urn:btih:" + hash + "&dn=" + dn, so
}

func GetPropertiesFromStream(stream Stream) Properties {
    var properties Properties

    titleSplit := strings.Split(stream.Title, "\n")
    emojiString := ""
    for _, title := range titleSplit {
        if strings.Contains(title, "👤") && strings.Contains(title, "💾") && strings.Contains(title, "⚙️") {
            emojiString = title
        }
    }

    emojiValues := extractEmojiValues(emojiString)
    magnet, files := parseLink(stream.Url)

    properties.Title = strings.ReplaceAll(titleSplit[0], " ", ".")
    properties.Link = magnet
    properties.Seeds = emojiValues[0]
    properties.Size = emojiValues[1]
    properties.Source = emojiValues[2]
    properties.Release = "[torrentio: " + properties.Source + "]"
    properties.Files = files

    return properties
}

func getByRegex(pattern string, streams []Stream) Stream {
    for _, stream := range streams {
        bingeGroup := strings.ToLower(stream.BehaviorHints.BingeGroup)

        matched, err := regexp.MatchString(pattern, bingeGroup)
        if err != nil {
            fmt.Printf("Error matching regular expression: %v\n", err)
            return Stream{}
        }

        if matched {
            return stream
        }
    }

    return Stream{}
}

type Filter struct {
    include bool
    pattern string
}

func getByFilters(filter []Filter, streams []Stream) Stream {
    for _, stream := range streams {
        match := true
        for _, filter := range filter {
            if filter.include {
                matched, err := regexp.MatchString(filter.pattern, stream.Title)
                if err != nil {
                    fmt.Printf("Error matching regular expression: %v\n", err)
                    return Stream{}
                }

                if !matched {
                    match = false
                    break
                }
            }

            if !filter.include {
                matched, err := regexp.MatchString(filter.pattern, stream.Title)
                if err != nil {
                    fmt.Printf("Error matching regular expression: %v\n", err)
                    return Stream{}
                }

                if matched {
                    match = false
                    break
                }
            }
        }

        if match {
            return stream
        }
    }

    return Stream{}
}

func FilterFormats(streams []Stream) []Stream {
    var results []Stream

    excludeBadEpisodeListing := Filter{
        include: false,
        pattern: `\.\d{1,2}x\d{1,2}\.`,
    }

    excludeHdr := Filter{
        include: false,
        pattern: `hdr`,
    }

    filters := [][]Filter{
        // includes 4k and HDR
        {
            {
                include: true,
                pattern: `2160p.*hdr|hdr.*2160p|4k.*hdr|hdr.*4k`,
            },
            excludeBadEpisodeListing,
        },
        // includes 4k and not HDR
        {
            {
                include: true,
                pattern: `2160p|4k`,
            },
            excludeHdr,
            excludeBadEpisodeListing,
        },
        // includes 1080p and HDR
        {
            {
                include: true,
                pattern: `1080p.*hdr|hdr.*1080p`,
            },
            excludeBadEpisodeListing,
        },
        // includes 1080p and not HDR
        {
            {
                include: true,
                pattern: `1080p`,
            },
            excludeHdr,
            excludeBadEpisodeListing,
        },
    }

    for _, filter := range filters {
        result := getByFilters(filter, streams)
        if result != (Stream{}) {
            results = append(results, result)
        }
    }

    if len(results) == 0 {
        results = append(results, streams[0])
    }

    return results
}
