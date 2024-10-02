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

func (h *Handler) Execute(filter *Filter) ([]song.Song, error) {
	return h.songRepo.FindMatching(func(song *song.Song) bool {
		if filter.GroupName != "" && song.GroupName != filter.GroupName {
			return false
		}

		if filter.SongName != "" && song.SongName != filter.SongName {
			return false
		}

		if filter.ReleaseDate != "" && song.ReleaseDate != filter.ReleaseDate {
			return false
		}

		if filter.Text != "" && song.Text != filter.Text {
			return false
		}

		if filter.Link != "" && song.Link != filter.Link {
			return false
		}

		return true
	})
}
