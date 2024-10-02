package query

import (
	"errors"
	"strings"

	"github.com/mishakrpv/musiclib/internal/domain/song"
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
		return nil, err
	}

	verses := strings.Split(song.Text, "\n\n")
	if page > 0 && len(verses) >= page {
		return &verses[page-1], nil
	}

	return nil, errors.New("verse not found")
}
