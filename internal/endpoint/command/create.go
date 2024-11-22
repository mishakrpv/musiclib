package command

import (
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/pkg/infra/musicinfo"
	"github.com/rs/zerolog/log"
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
	musicinfoClient musicinfo.Client
}

func NewCreateCommand(repo song.Repository,
	musicinfoClient musicinfo.Client) *CreateCommand {
	return &CreateCommand{songRepo: repo, musicinfoClient: musicinfoClient}
}

func (h *CreateCommand) Execute(request *CreateRequest) (*CreateResponse, error) {
	songDetail, err := h.musicinfoClient.GetSongDetail(request.Group, request.Song)
	if err != nil {
		log.Error().Err(err).Msg("An error occured getting SongDetail")
		return nil, err
	}

	log.Debug().
		Str("link", songDetail.Link).
		Str("text", songDetail.Text).
		Str("releaseDate", songDetail.ReleaseDate).
		Msg("SongDetail retrieved from API successfully")

	song := song.NewSong(request.Group, request.Song,
		songDetail.ReleaseDate, songDetail.Text, songDetail.Link)

	err = h.songRepo.Create(song)
	if err != nil {
		log.Error().Err(err).Msg("An error occured creating song")
		return nil, err
	}

	log.Debug().Str("id", song.Id.String()).Msg("Song created successfully")

	return &CreateResponse{Song: song}, nil
}
