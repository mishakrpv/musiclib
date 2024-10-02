package create

import (
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/internal/infrastructure/services/clients"

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

	song := song.NewSong(request.Group, request.Song,
		songDetail.ReleaseDate, songDetail.Text, songDetail.Link)

	err = h.songRepo.Create(song)
	if err != nil {
		zap.L().Error("An error occured creating song", zap.Error(err))
		return nil, err
	}

	return &Response{Song: song}, nil
}
