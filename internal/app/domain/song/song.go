package song

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	Id          uuid.UUID `json:"id"`
	GroupName   string    `json:"group"`
	SongName    string    `json:"song"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

func NewSong(
	groupName string,
	songName string,
	releaseDate time.Time,
	text string,
	link string,
) *Song {
	return &Song{
		Id:          uuid.New(),
		GroupName:   groupName,
		SongName:    songName,
		ReleaseDate: releaseDate,
		Text:        text,
		Link:        link,
	}
}

func (Song) TableName() string {
	return "songs"
}
