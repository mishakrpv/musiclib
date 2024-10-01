package postgres

import (
	"fmt"
	"os"

	"github.com/mishakrpv/musiclib/internal/app/domain/song"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	user     = os.Getenv("DB_USER")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository() song.Repository {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, database, port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Cannot open database connection", zap.Error(err))
	}

	zap.L().Info("Database connection has been opened", zap.String("dsn", dsn))

	return &SongRepository{
		db: db,
	}
}

func (s *SongRepository) Create(song *song.Song) error {
	panic("unimplemented")
}

func (s *SongRepository) Delete(group string, song string) error {
	panic("unimplemented")
}

func (s *SongRepository) FindMatching(predicate func(song *song.Song) bool) ([]song.Song, error) {
	panic("unimplemented")
}

func (s *SongRepository) Get(group string, song string) (*song.Song, error) {
	panic("unimplemented")
}

func (s *SongRepository) Update(song *song.Song) error {
	panic("unimplemented")
}
