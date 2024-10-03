package query

import "github.com/mishakrpv/musiclib/internal/domain/song"

type Filter struct {
	GroupName   string `form:"group" maxLength:"255"`
	SongName    string `form:"song" maxLength:"255"`
	ReleaseDate string `form:"date" maxLength:"10"`
	Text        string `form:"text"`
	Link        string `form:"link" maxLength:"255"`
}

type SongsQuery struct {
	songRepo song.Repository
}

func NewSongsQuery(songRepo song.Repository) *SongsQuery {
	return &SongsQuery{
		songRepo: songRepo,
	}
}

func (h *SongsQuery) Execute(filter *Filter) ([]*song.Song, error) {
	predicate := &song.Song{
		GroupName:   filter.GroupName,
		SongName:    filter.SongName,
		ReleaseDate: filter.ReleaseDate,
		Text:        filter.Text,
		Link:        filter.Link,
	}

	return h.songRepo.FindMatching(predicate)
}
