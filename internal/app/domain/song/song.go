package song

import (
	"time"
)

type Song struct {
	Group       string    `json:"group"`
	SongName    string    `json:"song"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

func NewSong(
	group string,
	songName string,
	releaseDate time.Time,
	text string,
	link string,
) *Song {
	return &Song{
		Group:       group,
		SongName:    songName,
		ReleaseDate: releaseDate,
		Text:        text,
		Link:        link,
	}
}
