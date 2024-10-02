package seed

import (
	"github.com/mishakrpv/musiclib/internal/domain/song"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type seed struct {
	Run func(*gorm.DB) error
}

func Seed(db *gorm.DB) {
	zap.L().Info("Start seeding")
	for _, seed := range all() {
		if err := seed.Run(db); err != nil {
			zap.L().Error("An error occured seeding db", zap.Error(err))
		}
	}
	zap.L().Info("Db seeded")
}

func all() []seed {
	return []seed{
		{
			Run: func(db *gorm.DB) error {
				return createSong(db, "Prince", "Soft and Wet")
			},
		},
		{
			Run: func(db *gorm.DB) error {
				return createSong(db, "The Beatles", "Ticket to Ride")
			},
		},
		{
			Run: func(db *gorm.DB) error {
				return createSong(db, "Black Sabbath", "Iron Man")
			},
		},
		{
			Run: func(db *gorm.DB) error {
				return createSong(db, "Black Sabbath", "Paranoid")
			},
		},
	}
}

func createSong(db *gorm.DB, groupName string, songName string) error {
	return db.Create(&song.Song{Id: uuid.New(), GroupName: groupName, SongName: songName}).Error
}
