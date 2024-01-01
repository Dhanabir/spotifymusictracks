package models

type MusicTracks struct {
	Items []MusicTrack
}

type MusicTrack struct {
	Name       string          `json:"name"`
	Popularity int             `json:"popularity"`
	ExtID      ExternalID      `json:"external_ids"`
	Artists    []ArtistDetails `json:"artists"`
	Album      Album           `json:"album"`
}

type ExternalID struct {
	ISRC string `json:"isrc"`
}

type ArtistDetails struct {
	Name string `json:"name"`
}

type Album struct {
	Images []ImageDetails `json:"images"`
}

type ImageDetails struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
