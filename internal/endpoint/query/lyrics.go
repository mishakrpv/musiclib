package query

import (
	"strings"

	"github.com/mishakrpv/musiclib/internal/apperror"
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/rs/zerolog/log"
)

type LyricsQuery struct {
	songRepo song.Repository
}

func NewLyricsQuery(songRepo song.Repository) *LyricsQuery {
	return &LyricsQuery{
		songRepo: songRepo,
	}
}

func (h *LyricsQuery) Execute(id string, page int) (*string, error) {
	song, err := h.songRepo.Get(id)
	if err != nil {
		return nil, apperror.ErrSongNotFound
	}

	log.Debug().Str("id", song.Id.String()).Msg("Song retrieved from db successfull")

	verses := strings.Split(song.Text, "\n\n")
	if page > 0 && len(verses) >= page {
		return &verses[page-1], nil
	}

	return nil, apperror.ErrVerseNotFound
}
