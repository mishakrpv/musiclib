package mock

import "github.com/mishakrpv/musiclib/internal/app/domain/song"

type SongRepository struct {
}

// Create implements song.Repository.
func (s *SongRepository) Create(song *song.Song) error {
	return nil
}

// Delete implements song.Repository.
func (s *SongRepository) Delete(group string, song string) error {
	return nil
}

// FindMatching implements song.Repository.
func (s *SongRepository) FindMatching(predicate func(song *song.Song) bool) ([]song.Song, error) {
	panic("unimplemented")
}

// Get implements song.Repository.
func (s *SongRepository) Get(group string, song string) (*song.Song, error) {
	panic("unimplemented")
}

// Update implements song.Repository.
func (s *SongRepository) Update(song *song.Song) error {
	panic("unimplemented")
}

func NewSongRepository() song.Repository {
	return &SongRepository{}
}
