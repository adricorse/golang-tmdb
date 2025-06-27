package tmdb

import "fmt"

// FindByID type is a struct for find JSON response.
type FindByID struct {
	MovieResults     []MovieMedia    `json:"movie_results,omitempty"`
	PersonResults    []PersonResults `json:"person_results"`
	TvResults        []TVShowMedia   `json:"tv_results,omitempty"`
	TvEpisodeResults []struct {
		AirDate        string `json:"air_date"`
		EpisodeNumber  int    `json:"episode_number"`
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Overview       string `json:"overview"`
		ProductionCode string `json:"production_code"`
		SeasonNumber   int    `json:"season_number"`
		ShowID         int64  `json:"show_id"`
		StillPath      string `json:"still_path"`
		VoteMetrics
	} `json:"tv_episode_results,omitempty"`
	TvSeasonResults []struct {
		AirDate      string `json:"air_date"`
		Name         string `json:"name"`
		ID           int64  `json:"id"`
		SeasonNumber int    `json:"season_number"`
		ShowID       int64  `json:"show_id"`
	} `json:"tv_season_results,omitempty"`
}

// GetFindByID the find method makes it easy to search for objects in our
// database by an external id. For example, an IMDB ID.
//
// This method will search all objects (movies, TV shows and people)
// and return the results in a single response.
//
// https://developers.themoviedb.org/3/find/find-by-id
func (c *Client) GetFindByID(
	id string,
	urlOptions map[string]string,
) (*FindByID, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s/find/%s?api_key=%s%s",
		baseURL, id, c.apiKey, options,
	)
	findByID := FindByID{}
	if err := c.get(tmdbURL, &findByID); err != nil {
		return nil, err
	}
	return &findByID, nil
}
