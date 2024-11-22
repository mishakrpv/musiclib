package db

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	migrationsPath := "migrations"

	var files []fs.FileInfo
	err := filepath.WalkDir(migrationsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := d.Info()
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() && filepath.Ext(d.Name()) == ".sql" {
			files = append(files, fileInfo)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, file := range files {
		sqlScript, err := os.ReadFile(filepath.Join(migrationsPath, file.Name()))
		if err != nil {
			log.Error().Err(err).Str("from_filename", file.Name()).Msg("Failed to read from")
			return err
		}

		err = db.Exec(string(sqlScript)).Error
		if err != nil {
			log.Error().Err(err).Str("from_filename", file.Name()).Msg("Failed to execute script")
			return err
		}

		log.Info().
			Str("from_filename", file.Name()).
			Msg("Migration has been applied")
	}

	log.Info().
		Msg("All migrations has been applied successfully")

	return nil
}
