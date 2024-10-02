package query

import "github.com/mishakrpv/musiclib/internal/domain/song"

type Filter struct {
	GroupName   string `form:"group"`
	SongName    string `form:"song"`
	ReleaseDate string `form:"date"`
	Text        string `form:"text"`
	Link        string `form:"link"`
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
		GroupName: filter.GroupName,
		SongName: filter.SongName,
		ReleaseDate: filter.ReleaseDate,
		Text: filter.Text,
		Link: filter.Link,
	}

	return h.songRepo.FindMatching(predicate)
}
