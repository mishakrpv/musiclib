package command

import (
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/internal/infrastructure/service"

	"go.uber.org/zap"
)

type CreateRequest struct {
	Group string `json:"group" binding:"required" maxLength:"255"`
	Song  string `json:"song" binding:"required" maxLength:"255"`
}

type CreateResponse struct {
	Song *song.Song
}

type CreateCommand struct {
	songRepo        song.Repository
	musicInfoClient service.MusicInfoClient
}

func NewCreateCommand(repo song.Repository,
	musicInfoClient service.MusicInfoClient) *CreateCommand {
	return &CreateCommand{songRepo: repo, musicInfoClient: musicInfoClient}
}

func (h *CreateCommand) Execute(request *CreateRequest) (*CreateResponse, error) {
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

	return &CreateResponse{Song: song}, nil
}
