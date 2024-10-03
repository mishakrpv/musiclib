package gorm

import (
	"io/fs"
	"os"
	"path/filepath"

	"go.uber.org/zap"
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
			zap.L().Fatal("Failed to read from", zap.String("filename", file.Name()), zap.Error(err))
			return err
		}

		err = db.Exec(string(sqlScript)).Error
		if err != nil {
			zap.L().Fatal("Failed to execute script", zap.String("from_filename", file.Name()), zap.Error(err))
			return err
		}

		zap.L().Info("Migration has been applied", zap.String("from_filename", file.Name()))
	}

	zap.L().Info("All migrations has been applied successfully")

	return nil
}
