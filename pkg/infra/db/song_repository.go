package db

import (
	"context"
	"fmt"

	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/pkg/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(ctx context.Context, cfg *config.DBConfig) (song.Repository, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.HOST, cfg.User, cfg.Pwd, cfg.Database, cfg.Port)

	log.Info().
		Str("dsn", dsn).
		Msg("Postgres dsn read from configurations")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Cannot open database connection")
		return nil, err
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		logger := log.Ctx(ctx)
		logger.Info().Msg("Closing DB connection...")
		sqlDB, err := db.DB()
		if err != nil {
			log.Error().Err(err).Msg("DB conn forced to close with error")
		} else {
			sqlDB.Close()
		}
	}(ctx)

	log.Info().
		Str("dsn", dsn).
		Msg("Database connection has been opened")

	log.Info().
		Msg("Start migrating")

	err = Migrate(db)
	if err != nil {

		log.Error().Err(err).Msg("An error occured migrating database")
		return nil, err
	}
	log.Info().
		Msg("Db migrated")

	return &SongRepository{
		db: db,
	}, nil
}

func (repo *SongRepository) Create(song *song.Song) error {
	return repo.db.Create(song).Error
}

func (repo *SongRepository) Delete(id string) error {
	return repo.db.Where("id= ?", id).Delete(&song.Song{}).Error
}

func (repo *SongRepository) FindMatching(predicate *song.Song) ([]*song.Song, error) {
	var songs []*song.Song
	err := repo.db.Where(*predicate).Find(&songs).Error
	return songs, err
}

func (repo *SongRepository) Get(id string) (*song.Song, error) {
	var song = &song.Song{}
	err := repo.db.First(&song, "id = ?", id).Error
	return song, err
}

func (repo *SongRepository) Update(song *song.Song) error {
	return repo.db.Model(&song).Updates(*song).Error
}
