package models

type Episode struct {
    ID            int    `json:"id"`
    SeriesID      int    `json:"series_id"`
    SeasonNumber  int    `json:"season_number"`
    EpisodeNumber int    `json:"episode_number"`
    Title         string `json:"title"`
    Description   string `json:"description"`
    Duration      int    `json:"duration"`
    ThumbnailURL  string `json:"thumbnail_url"`
    EpisodeURL    string `json:"episode"`       // stored as varchar(255)
    ReleaseDate   string `json:"release_date"`  // use string (YYYY-MM-DD)
}
