package tmdb

import (
	"fmt"

	json "github.com/goccy/go-json"
)

func (c *CreditsDetails) UnmarshalJSON(data []byte) error {
	type alias CreditsDetails
	var raw struct {
		alias
		Media json.RawMessage `json:"media"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*c = CreditsDetails(raw.alias)
	var mediaType struct {
		MediaType string `json:"media_type"`
	}
	if err := json.Unmarshal(raw.Media, &mediaType); err != nil {
		return err
	}

	switch mediaType.MediaType {
	case "movie":
		var movie CreditMovieMedia
		if err := json.Unmarshal(raw.Media, &movie); err != nil {
			return err
		}
		c.Media = movie
	case "tv":
		var tv CreditTVShowMedia
		if err := json.Unmarshal(raw.Media, &tv); err != nil {
			return err
		}
		c.Media = tv
	default:
		return fmt.Errorf("unknown media_type: %s", mediaType.MediaType)
	}
	return nil
}

func (p *PersonResults) UnmarshalJSON(data []byte) error {
	type alias PersonResults
	var raw struct {
		alias
		KnownFor []json.RawMessage `json:"known_for"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*p = PersonResults(raw.alias)
	for _, item := range raw.KnownFor {
		var mediaType struct {
			MediaType string `json:"media_type"`
		}
		if err := json.Unmarshal(item, &mediaType); err != nil {
			return err
		}

		switch mediaType.MediaType {
		case "movie":
			var movie MovieMedia
			if err := json.Unmarshal(item, &movie); err != nil {
				return err
			}
			p.KnownFor = append(p.KnownFor, movie)
		case "tv":
			var tv TVShowMedia
			if err := json.Unmarshal(item, &tv); err != nil {
				return err
			}
			p.KnownFor = append(p.KnownFor, tv)
		default:
			return fmt.Errorf("unknown media_type: %s", mediaType.MediaType)
		}
	}
	return nil
}
