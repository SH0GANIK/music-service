package model

type Song struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	SongDetails `json:"songDetails"`
}

type SongDetails struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type PaginatedSongsResponse struct {
	Data     []*Song `json:"data"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
}

type PaginatedSongTextResponse struct {
	Data     []string `json:"data"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
}
