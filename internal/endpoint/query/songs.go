package query

import "github.com/mishakrpv/musiclib/internal/domain/song"

type Filter struct {
	GroupName   string `form:"froup"`
	SongName    string `form:"song"`
	ReleaseDate string `form:"date"`
	Text        string `form:"text"`
	Link        string `form:"link"`
}

type Handler struct {
	songRepo song.Repository
}

func NewHandler(songRepo song.Repository) *Handler {
	return &Handler{
		songRepo: songRepo,
	}
}

func (h *Handler) Execute(filter *Filter) ([]*song.Song, error) {
	predicate := &song.Song{
		GroupName: filter.GroupName,
		SongName: filter.SongName,
		ReleaseDate: filter.ReleaseDate,
		Text: filter.Text,
		Link: filter.Link,
	}

	return h.songRepo.FindMatching(predicate)
}
