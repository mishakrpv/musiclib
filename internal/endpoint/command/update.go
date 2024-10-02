package command

type UpdateRequest struct {
	GroupName   string `json:"group" binding:"required"`
	SongName    string `json:"song" binding:"required"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
