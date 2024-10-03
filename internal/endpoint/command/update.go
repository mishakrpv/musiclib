package command

type UpdateRequest struct {
	GroupName   string `json:"group" binding:"required" maxLength:"255"`
	SongName    string `json:"song" binding:"required" maxLength:"255"`
	ReleaseDate string `json:"release_date" maxLength:"10"`
	Text        string `json:"text"`
	Link        string `json:"link" maxLength:"255"`
}
