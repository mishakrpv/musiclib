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

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
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
	zap.L().Info("Start handling request")

	filter := &query.Filter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		zap.L().Error("An error occured binding query", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page := c.GetInt("page")
	amount := c.GetInt("songs")

	zap.L().Debug("Query bound",
		zap.String("group", filter.GroupName),
		zap.String("song", filter.SongName),
		zap.String("date", filter.ReleaseDate),
		zap.String("text", filter.Text),
		zap.String("link", filter.Link),
		zap.Int("page", page),
		zap.Int("amount", amount))

	amount--
	page--

	link, err := url.QueryUnescape(filter.Link)
	if err != nil {
		zap.L().Error("An error occured decoding link", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter.Link = link
	zap.L().Debug("Query link decoded", zap.String("link", filter.Link))

	qr := query.NewSongsQuery(r.songRepo)

	response, err := qr.Execute(filter)
	if err != nil {
		zap.L().Error("Something went wrong while executing query", zap.Error(err))
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

	zap.L().Info("Request handled successfully")
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
	zap.L().Info("Start handling request")

	id := c.Param("song_id")
	_, err := uuid.Parse(id)
	if err != nil {
		zap.L().Error("An error occured parsing id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
		return
	}

	page := c.GetInt("page")
	zap.L().Debug("Params bound", zap.String("id", id), zap.Int("page", page))

	query := query.NewLyricsQuery(r.songRepo)

	verse, err := query.Execute(id, page)
	if err != nil {
		var status int = http.StatusInternalServerError
		if errors.Is(err, apperror.ErrVerseNotFound) || errors.Is(err, apperror.ErrSongNotFound) {
			status = http.StatusNotFound
		}
		zap.L().Error("An error occured executing query", zap.Error(err))
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	zap.L().Info("Request handled successfully")
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
	zap.L().Info("Start handling request")

	id := c.Param("song_id")
	_, err := uuid.Parse(id)
	if err != nil {
		zap.L().Error("An error occured parsing id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
		return
	}

	zap.L().Debug("Param bound", zap.String("id", id))

	err = r.songRepo.Delete(id)
	if err != nil {
		zap.L().Error("Something went wrong while deleting song", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	zap.L().Info("Request handled successfully")
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
	zap.L().Info("Start handling request")

	songId := c.Param("song_id")
	id, err := uuid.Parse(songId)
	if err != nil {
		zap.L().Error("An error occured parsing id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
		return
	}

	zap.L().Debug("Param bound", zap.String("id", songId))

	request := &command.UpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.L().Error("An error occured binding request", zap.Error(err))
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
		zap.L().Error("Something went wrong while updating song", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	zap.L().Info("Request handled successfully")
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
		zap.L().Error("An error occured binding request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zap.L().Debug("Body bound",
		zap.String("group", request.Group),
		zap.String("song", request.Song))

	response, err := cmd.Execute(request)
	if err != nil {
		var status = http.StatusInternalServerError
		if errors.Is(err, apperror.ErrSongNotFound) {
			status = http.StatusNotFound
		}
		zap.L().Error("An error occured executing command", zap.Error(err))
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	zap.L().Info("Request handled successfully")
	c.JSON(http.StatusOK, &response)
}
