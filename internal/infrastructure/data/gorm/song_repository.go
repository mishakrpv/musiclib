package gorm

import (
	"fmt"
	"os"

	"github.com/mishakrpv/musiclib/internal/domain/song"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	database = os.Getenv("DB_DATABASE")
	port     = os.Getenv("DB_PORT")
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository() song.Repository {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, database, port)

	zap.L().Info("Postgres dsn read from configurations", zap.String("dsn", dsn))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Cannot open database connection", zap.Error(err))
	}

	zap.L().Info("Database connection has been opened", zap.String("dsn", dsn))

	zap.L().Info("Start migrating")
	err = Migrate(db)
	if err != nil {
		zap.L().Fatal("An error occured migrating database", zap.Error(err))
	}
	zap.L().Info("Db migrated")

	return &SongRepository{
		db: db,
	}
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
