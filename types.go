package tmdb

// BelongsToCollection represents information about a collection to which a movie belongs,
// including its unique identifier, name, poster image path, and backdrop image path.
type BelongsToCollection struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

// ProductionCompany represents a company involved in the production of a movie or TV show.
// It includes the company's name, unique identifier, logo path, and country of origin.
type ProductionCompany struct {
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

// Genre represents a movie or TV show genre with its unique identifier and name.
type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ProductionCountry represents a country involved in the production of a movie or TV show,
// containing its ISO 3166-1 code and the country's name.
type ProductionCountry struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

// SpokenLanguage represents a language spoken in a media item, including its ISO 639-1 code and name.
type SpokenLanguage struct {
	Iso639_1 string `json:"iso_639_1"`
	Name     string `json:"name"`
}

// TranslationData represents the translated information for a media item,
// including its title, name, overview, homepage, runtime, and tagline.
// Fields are mapped to their respective JSON keys and may be omitted if empty.
type TranslationData struct {
	Title    string `json:"title,omitempty"`
	Name     string `json:"name,omitempty"`
	Overview string `json:"overview,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Runtime  *int   `json:"runtime,omitempty"`
	Tagline  string `json:"tagline,omitempty"`
}

// Translation represents a translation entry with language and country codes,
// localized and English names, and associated translation data.
type Translation struct {
	Iso3166_1   string          `json:"iso_3166_1"`
	Iso639_1    string          `json:"iso_639_1"`
	Name        string          `json:"name"`
	EnglishName string          `json:"english_name"`
	Data        TranslationData `json:"data"`
}

// PaginatedResultsMeta represents metadata for paginated API responses,
// including the current page, total number of results, and total number of pages.
type PaginatedResultsMeta struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
}

// Season represents a TV show season with details such as air date, episode count, name, overview, poster path, season number, vote average, and associated show ID.
type Season struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	VoteAverage  float32 `json:"vote_average"`
	AirDate      string  `json:"air_date"`
	SeasonNumber int     `json:"season_number"`
	EpisodeCount int     `json:"episode_count"`
}

// LastEpisodeToAir represents the details of the most recently aired episode of a TV show.
// It includes information such as air date, episode and season numbers, production code,
// episode overview, voting statistics, and related media paths.
type LastEpisodeToAir struct {
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
}

// NextEpisodeToAir represents the details of the next episode scheduled to air for a TV show.
// It includes information such as air date, episode and season numbers, show and episode IDs,
// episode name and overview, production code, still image path, and voting statistics.
type NextEpisodeToAir struct {
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
}

// Network represents a television network with its identifying information,
// including name, unique ID, logo path, and country of origin.
type Network struct {
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

// CreatedBy represents a person who created a particular media item, such as a TV show or movie.
// It includes identifying information such as ID, credit ID, name, gender, and profile image path.
type CreatedBy struct {
	ID          int64  `json:"id"`
	CreditID    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type VoteMetrics struct {
	VoteCount   int64   `json:"vote_count"`
	VoteAverage float32 `json:"vote_average"`
}

type ImageBase struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	Width       int     `json:"width"`
	VoteMetrics
}

type PersonBase struct {
	Adult              bool    `json:"adult"`
	ID                 int64   `json:"id"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float32 `json:"popularity"`
	Gender             int     `json:"gender"`
	KnownForDepartment string  `json:"known_for_department"`
	ProfilePath        string  `json:"profile_path"`
}

type Person struct {
	PersonBase
}

type PersonResults struct {
	PersonBase
	MediaType string  `json:"media_type"`
	KnownFor  []Media `json:"known_for"`
}

type Media interface {
	GetMediaType() string
}

func (movie MovieMedia) GetMediaType() string { return movie.MediaType }
func (tv TVShowMedia) GetMediaType() string   { return tv.MediaType }

type CreditMedia interface {
	GetMediaType() string
}

func (movie CreditMovieMedia) GetMediaType() string { return movie.MediaType }
func (tv CreditTVShowMedia) GetMediaType() string   { return tv.MediaType }

type CreditTVShowMedia struct {
	TVShowMedia
	Character string `json:"character"`
}

type CreditMovieMedia struct {
	MovieMedia
	Character string `json:"character"`
}

type Episode struct {
	AirDate       string `json:"air_date"`
	EpisodeNumber int64  `json:"episode_number"`
	Name          string `json:"name"`
	Overview      string `json:"overview"`
	SeasonNumber  int    `json:"season_number"`
	StillPath     string `json:"still_path"`
}

type VideoBase struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	ID               int64   `json:"id"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	Popularity       float32 `json:"popularity"`
	VoteMetrics
}

type MovieMeta struct {
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	ReleaseDate   string `json:"release_date"`
	Video         bool   `json:"video"`
}

type Movie struct {
	VideoBase
	MovieMeta
	GenreIDs []int64 `json:"genre_ids"`
}

type MovieMedia struct {
	Movie
	MediaType string `json:"media_type"`
}

type TVShowMeta struct {
	Name          string         `json:"name"`
	OriginalName  string         `json:"original_name"`
	FirstAirDate  string         `json:"first_air_date"`
	OriginCountry []string       `json:"origin_country"`
	Episodes      []Episode      `json:"episodes"`
	Seasons       []CreditSeason `json:"seasons"`
}

type TVShow struct {
	VideoBase
	TVShowMeta
	GenreIDs []int64 `json:"genre_ids"`
}

type TVShowMedia struct {
	TVShow
	MediaType string `json:"media_type"`
}

type VideoDetails struct {
	VideoBase
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
}

type MovieDetailsItem struct {
	VideoDetails
	MovieMeta
	OriginCountry       []string            `json:"origin_country"`
	BelongsToCollection BelongsToCollection `json:"belongs_to_collection"`
	Budget              int64               `json:"budget"`
	IMDbID              string              `json:"imdb_id"`
	Revenue             int64               `json:"revenue"`
	Runtime             int                 `json:"runtime"`
}

type TVShowDetailsItem struct {
	VideoDetails
	TVShowMeta
	CreatedBy        []CreatedBy      `json:"created_by"`
	EpisodeRunTime   []int            `json:"episode_run_time"`
	InProduction     bool             `json:"in_production"`
	Languages        []string         `json:"languages"`
	LastAirDate      string           `json:"last_air_date"`
	LastEpisodeToAir LastEpisodeToAir `json:"last_episode_to_air"`
	NextEpisodeToAir NextEpisodeToAir `json:"next_episode_to_air"`
	Networks         []Network        `json:"networks"`
	NumberOfEpisodes int              `json:"number_of_episodes"`
	NumberOfSeasons  int              `json:"number_of_seasons"`
	Type             string           `json:"type"`
}
