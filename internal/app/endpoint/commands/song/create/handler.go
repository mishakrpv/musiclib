package create

import (
	"time"

	"github.com/mishakrpv/musiclib/internal/app/domain/song"
	"github.com/mishakrpv/musiclib/internal/app/infrastructure/services/clients"

	"go.uber.org/zap"
)

type Handler struct {
	songRepo        song.Repository
	musicInfoClient clients.MusicInfoClient
}

func NewHandler(repo song.Repository,
	musicInfoClient clients.MusicInfoClient) *Handler {
	return &Handler{songRepo: repo, musicInfoClient: musicInfoClient}
}

func (h *Handler) Execute(request *Request) (*Response, error) {
	songDetail, err := h.musicInfoClient.GetSongDetail(request.Group, request.Song)
	if err != nil {
		zap.L().Error("An error occured getting SongDetail", zap.Error(err))
		return nil, err
	}

	format := "2006.01.02"
	date, err := time.Parse(format, songDetail.ReleaseDate)
	if err != nil {
		zap.L().Error("An error occured parsing SongDetail release date", zap.Error(err))
		return nil, err
	}

	song := song.NewSong(request.Group, request.Song,
		date, songDetail.Text, songDetail.Link)

	err = h.songRepo.Create(song)
	if err != nil {
		zap.L().Error("An error occured creating song", zap.Error(err))
		return nil, err
	}

	return &Response{Song: song}, nil
}
