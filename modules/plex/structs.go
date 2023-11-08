package plex

import "encoding/xml"

type TvDirectory struct {
	XMLName xml.Name `xml:"Directory"`
	RatingKey string `xml:"ratingKey,attr"`
	Key string `xml:"key,attr"`
	ParentRatingKey string `xml:"parentRatingKey,attr"`
	Guid string `xml:"guid,attr"`
	ParentGuid string `xml:"parentGuid,attr"`
	ParentStudio string `xml:"parentStudio,attr"`
	Type string `xml:"type,attr"`
	Title string `xml:"title,attr"`
	ParentKey string `xml:"parentKey,attr"`
	ParentTitle string `xml:"parentTitle,attr"`
	Summary string `xml:"summary,attr"`
	Index string `xml:"index,attr"`
	ParentIndex string `xml:"parentIndex,attr"`
	ParentYear string `xml:"parentYear,attr"`
	Thumb string `xml:"thumb,attr"`
	Art string `xml:"art,attr"`
	ParentThumb string `xml:"parentThumb,attr"`
	ParentTheme string `xml:"parentTheme,attr"`
	LeafCount string `xml:"leafCount,attr"`
	ViewedLeafCount string `xml:"viewedLeafCount,attr"`
	AddedAt string `xml:"addedAt,attr"`
	UpdatedAt string `xml:"updatedAt,attr"`
}

type TvMetadata struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Size string `xml:"size,attr"`
	AllowSync string `xml:"allowSync,attr"`
	Art string `xml:"art,attr"`
	Identifier string `xml:"identifier,attr"`
	Key string `xml:"key,attr"`
	LibrarySectionID string `xml:"librarySectionID,attr"`
	LibrarySectionTitle string `xml:"librarySectionTitle,attr"`
	LibrarySectionUUID string `xml:"librarySectionUUID,attr"`
	MediaTagPrefix string `xml:"mediaTagPrefix,attr"`
	MediaTagVersion string `xml:"mediaTagVersion,attr"`
	NoCache string `xml:"nocache,attr"`
	ParentIndex string `xml:"parentIndex,attr"`
	ParentTitle string `xml:"parentTitle,attr"`
	ParentYear string `xml:"parentYear,attr"`
	Summary string `xml:"summary,attr"`
	Theme string `xml:"theme,attr"`
	Thumb string `xml:"thumb,attr"`
	Title1 string `xml:"title1,attr"`
	Title2 string `xml:"title2,attr"`
	ViewGroup string `xml:"viewGroup,attr"`
	ViewMode string `xml:"viewMode,attr"`
	Directory []TvDirectory `xml:"Directory"`
}

type Video struct {
	XMLName xml.Name `xml:"Video"`
	RatingKey string `xml:"ratingKey,attr"`
	Key string `xml:"key,attr"`
	ParentRatingKey string `xml:"parentRatingKey,attr"`
	GrandparentRatingKey string `xml:"grandparentRatingKey,attr"`
	Guid string `xml:"guid,attr"`
	ParentGuid string `xml:"parentGuid,attr"`
	GrandparentGuid string `xml:"grandparentGuid,attr"`
	Type string `xml:"type,attr"`
	Title string `xml:"title,attr"`
	TitleSort string `xml:"titleSort,attr"`
	GrandparentKey string `xml:"grandparentKey,attr"`
	ParentKey string `xml:"parentKey,attr"`
	GrandparentTitle string `xml:"grandparentTitle,attr"`
	ParentTitle string `xml:"parentTitle,attr"`
	ContentRating string `xml:"contentRating,attr"`
	Summary string `xml:"summary,attr"`
	Index string `xml:"index,attr"`
	ParentIndex string `xml:"parentIndex,attr"`
	Year string `xml:"year,attr"`
	Thumb string `xml:"thumb,attr"`
	Art string `xml:"art,attr"`
	ParentThumb string `xml:"parentThumb,attr"`
	GrandparentThumb string `xml:"grandparentThumb,attr"`
	GrandparentArt string `xml:"grandparentArt,attr"`
	GrandparentTheme string `xml:"grandparentTheme,attr"`
	Duration string `xml:"duration,attr"`
	OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
	AddedAt string `xml:"addedAt,attr"`
	UpdatedAt string `xml:"updatedAt,attr"`
	// Media Media `xml:"Media"`
}

type SeasonMetadata struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Size string `xml:"size,attr"`
	AllowSync string `xml:"allowSync,attr"`
	Art string `xml:"art,attr"`
	GrandparentContentRating string `xml:"grandparentContentRating,attr"`
	GrandparentRatingKey string `xml:"grandparentRatingKey,attr"`
	GrandparentStudio string `xml:"grandparentStudio,attr"`
	GrandparentTheme string `xml:"grandparentTheme,attr"`
	GrandparentThumb string `xml:"grandparentThumb,attr"`
	GrandparentTitle string `xml:"grandparentTitle,attr"`
	Identifier string `xml:"identifier,attr"`
	Key string `xml:"key,attr"`
	LibrarySectionID string `xml:"librarySectionID,attr"`
	LibrarySectionTitle string `xml:"librarySectionTitle,attr"`
	LibrarySectionUUID string `xml:"librarySectionUUID,attr"`
	MediaTagPrefix string `xml:"mediaTagPrefix,attr"`
	MediaTagVersion string `xml:"mediaTagVersion,attr"`
	NoCache string `xml:"nocache,attr"`
	ParentIndex string `xml:"parentIndex,attr"`
	ParentTitle string `xml:"parentTitle,attr"`
	SortAsc string `xml:"sortAsc,attr"`
	Summary string `xml:"summary,attr"`
	Theme string `xml:"theme,attr"`
	Thumb string `xml:"thumb,attr"`
	Title1 string `xml:"title1,attr"`
	Title2 string `xml:"title2,attr"`
	ViewGroup string `xml:"viewGroup,attr"`
	ViewMode string `xml:"viewMode,attr"`
	Video []Video `xml:"Video"`
}
