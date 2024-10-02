package lyrics

import (
	"errors"
	"strings"

	"github.com/mishakrpv/musiclib/internal/domain/song"
)

type Handler struct {
	songRepo song.Repository
}

func NewHandler(songRepo song.Repository) *Handler {
	return &Handler{
		songRepo: songRepo,
	}
}

func (h *Handler) Execute(id string, page int) (*string, error) {
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
