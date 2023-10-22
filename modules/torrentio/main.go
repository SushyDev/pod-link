package torrentio

import (
    "fmt"
    "regexp"
    "strings"
)

type Stream struct {
    Name          string `json:"name"`
    Title         string `json:"title"`
    InfoHash      string `json:"infoHash"`
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

    properties.Title = strings.ReplaceAll(titleSplit[0], " ", ".")
    properties.Link = "magnet:?xt=urn:btih:" + stream.InfoHash + "&dn=&tr="
    properties.Seeds = emojiValues[0]
    properties.Size = emojiValues[1]
    properties.Source = emojiValues[2]
    properties.Release = "[torrentio: " + properties.Source + "]"

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

func getByFilters(filters []Filter, streams []Stream) Stream {
    for _, stream := range streams {
        bingeGroup := strings.ToLower(stream.BehaviorHints.BingeGroup)

        match := true
        for _, filter := range filters {
            if filter.include {
                matched, err := regexp.MatchString(filter.pattern, bingeGroup)
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
                matched, err := regexp.MatchString(filter.pattern, bingeGroup)
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


func FilterResults(streams []Stream) []Stream {
    var results []Stream

    filters := [][]Filter{
        // includes 4k and HDR
        {
            {
                include: true,
                pattern: `4k.*hdr|hdr.*4k`,
            },
        },
        // includes 4k and not HDR
        {
            {
                include: true,
                pattern: `4k`,
            },
            {
                include: false,
                pattern: `hdr`,
            },
        },
        // includes 1080p and HDR
        {
            {
                include: true,
                pattern: `1080p.*hdr|hdr.*1080p`,
            },
        },
        // includes 1080p and not HDR
        {
            {
                include: true,
                pattern: `1080p`,
            },
            {
                include: false,
                pattern: `hdr`,
            },
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
