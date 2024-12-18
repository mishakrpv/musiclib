package router

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mishakrpv/musiclib/internal/apperror"
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/internal/endpoint/command"
	"github.com/mishakrpv/musiclib/internal/endpoint/query"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		Songs
// @Tags			query
// @Description	get all songs matching filters from the library
// @Produce		json
// @Param			group	query		string	false	"search by group"	maxlength(255)
// @Param			song	query		string	false	"search by song"	maxlength(255)
// @Param			date	query		string	false	"search by date"	maxlength(10)
// @Param			text	query		string	false	"search by text"
// @Param			link	query		string	false	"search by link"	maxlength(255)
// @Param			page	query		int		false	"page number"
// @Param			songs	query		int		false	"songs per page"	minimum(3)	maximum(150)
// @Success		200		{array}		song.Song
// @Failure		400		{string}	string	"error"
// @Failure		500		{string}	string	"error"
// @Router			/songs [get]
func (r *Router) SongsHandler(c *gin.Context) {
	log.Info().Msg("Start handling request")

	filter := &query.Filter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		log.Error().Err(err).Msg("An error occured binding query")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page := c.GetInt("page")
	amount := c.GetInt("songs")

	log.Debug().
		Str("group", filter.GroupName).
		Str("song", filter.SongName).
		Str("date", filter.ReleaseDate).
		Str("text", filter.Text).
		Str("link", filter.Link).
		Int("page", page).
		Int("amount", amount).
		Msg("Query bound")

	amount--
	page--

	link, err := url.QueryUnescape(filter.Link)
	if err != nil {
		log.Error().Err(err).Msg("An error occured decoding link")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter.Link = link
	log.Debug().Str("link", filter.Link).Msg("Query link decoded")

	qr := query.NewSongsQuery(r.songRepo)

	response, err := qr.Execute(filter)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong while executing query")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response = response[page*amount : (page+1)*amount]

	res := make([]interface{}, 0, len(response))
	for _, item := range response {
		if item != nil {
			res = append(res, item)
		}
	}

	log.Info().Msg("Request handled successfully")
	c.JSON(http.StatusOK, &res)
}

// @Summary		Lyrics
// @Tags			query
// @Description	get song's lyrics with pagination by verses
// @Produce		json
// @Param			song_id	path		string	true	"Song ID"	maxlength(255)
// @Param			page	query		int		false	"verse number"
// @Success		200		{string}	string	"verse"
// @Failure		400		{string}	string	"error"
// @Failure		404		{string}	string	"error"
// @Failure		500		{string}	string	"error"
// @Router			/songs/{song_id}/lyrics [get]
func (r *Router) LyricsHandler(c *gin.Context) {
	log.Info().Msg("Start handling request")

	id := c.Param("song_id")
	_, err := uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Msg("An error occured parsing id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
		return
	}

	page := c.GetInt("page")
	log.Debug().Str("id", id).Int("page", page).Msg("Params bound")

	query := query.NewLyricsQuery(r.songRepo)

	verse, err := query.Execute(id, page)
	if err != nil {
		var status = http.StatusInternalServerError
		if errors.Is(err, apperror.ErrVerseNotFound) || errors.Is(err, apperror.ErrSongNotFound) {
			status = http.StatusNotFound
		}
		log.Error().Err(err).Msg("An error occured executing query")
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	log.Info().Msg("Request handled successfully")
	c.JSON(http.StatusOK, gin.H{fmt.Sprintf("verse number %d:", page): *verse})
}

// @Summary		DeleteSong
// @Tags			command
// @Description	delete song
// @Param			song_id	path	string	true	"Song ID"	maxlength(255)
// @Success		200
// @Failure		400	{string}	string	"error"
// @Failure		500	{string}	string	"error"
// @Router			/songs/{song_id} [delete]
func (r *Router) DeleteSongHandler(c *gin.Context) {
	log.Info().Msg("Start handling request")

	id := c.Param("song_id")
	_, err := uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Msg("An error occured parsing id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
		return
	}

	log.Debug().Str("id", id).Msg("Param bound")

	err = r.songRepo.Delete(id)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong while deleting a song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info().Msg("Request handled successfully")
	c.Status(http.StatusOK)
}

// @Summary		UpdateSong
// @Tags			command
// @Description	update song
// @Accept			json
// @Param			song_id	path	string					true	"Song ID"	maxlength(255)
// @Param			request	body	command.UpdateRequest	true	"song to update"
// @Success		200
// @Failure		400	{string}	string	"error"
// @Failure		500	{string}	string	"error"
// @Router			/songs/{song_id} [put]
func (r *Router) UpdateSongHandler(c *gin.Context) {
	log.Info().Msg("Start handling request")

	songID := c.Param("song_id")
	id, err := uuid.Parse(songID)
	if err != nil {
		log.Error().Err(err).Msg("An error occured parsing id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
		return
	}

	log.Debug().Str("id", songID).Msg("Param bound")

	request := &command.UpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("An error occured binding request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = r.songRepo.Update(&song.Song{
		Id:          id,
		GroupName:   request.GroupName,
		SongName:    request.SongName,
		ReleaseDate: request.ReleaseDate,
		Text:        request.Text,
		Link:        request.Link,
	})
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong while updating song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info().Msg("Request handled successfully")
	c.Status(http.StatusOK)
}

// @Summary		CreateSong
// @Tags			command
// @Description	create song
// @Accept			json
// @Produce		json
// @Param			request	body		command.CreateRequest	true	"song to create"
// @Success		200		{object}	song.Song
// @Failure		400		{string}	string	"error"
// @Failure		404		{string}	string	"error"
// @Failure		500		{string}	string	"error"
// @Router			/songs [post]
func (r *Router) CreateSongHandler(c *gin.Context) {
	cmd := command.NewCreateCommand(r.songRepo, r.musicinfoClient)

	request := &command.CreateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("An error occured binding request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Debug().
		Str("group", request.Group).
		Str("song", request.Group).
		Msg("Body bound")

	response, err := cmd.Execute(request)
	if err != nil {
		var status = http.StatusInternalServerError
		if errors.Is(err, apperror.ErrSongNotFound) {
			status = http.StatusNotFound
		}
		log.Error().Err(err).Msg("An error occured executing command")
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	log.Info().Msg("Request handled successfully")
	c.JSON(http.StatusOK, &response)
}
