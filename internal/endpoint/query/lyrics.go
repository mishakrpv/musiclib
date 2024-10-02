package query

import (
	"strings"

	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/internal/apperror"
)

type LyricsHandler struct {
	songRepo song.Repository
}

func NewLyricsHandler(songRepo song.Repository) *LyricsHandler {
	return &LyricsHandler{
		songRepo: songRepo,
	}
}

func (h *LyricsHandler) Execute(id string, page int) (*string, error) {
	song, err := h.songRepo.Get(id)
	if err != nil {
		return nil, apperror.ErrSongNotFound
	}

	verses := strings.Split(song.Text, "\n\n")
	if page > 0 && len(verses) >= page {
		return &verses[page-1], nil
	}

	return nil, apperror.ErrVerseNotFound
}
