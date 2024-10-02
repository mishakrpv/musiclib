package postgres

import (
	"github.com/google/uuid"
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"gorm.io/gorm"
)

type Seed struct {
	Run func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		{
			Run: func(db *gorm.DB) error {
				return CreateSong(db, "Prince", "Soft and Wet")
			},
		},
		{
			Run: func(db *gorm.DB) error {
				return CreateSong(db, "The Beatles", "Ticket to Ride")
			},
		},
		{
			Run: func(db *gorm.DB) error {
				return CreateSong(db, "Black Sabbath", "Iron Man")
			},
		},
		{
			Run: func(db *gorm.DB) error {
				return CreateSong(db, "Black Sabbath", "Paranoid")
			},
		},
	}
}

func CreateSong(db *gorm.DB, groupName string, songName string) error {
	return db.Create(&song.Song{Id: uuid.New(), GroupName: groupName, SongName: songName}).Error
}
