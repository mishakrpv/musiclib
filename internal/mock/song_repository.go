package mock

import "github.com/mishakrpv/musiclib/internal/domain/song"

type SongRepository struct {
}

func NewSongRepository() song.Repository {
	return &SongRepository{}
}

func (s *SongRepository) Create(song *song.Song) error {
	return nil
}

func (s *SongRepository) Delete(group string, song string) error {
	panic("unimplemented")
}

func (s *SongRepository) FindMatching(predicate *song.Song) ([]*song.Song, error) {
	panic("unimplemented")
}

func (s *SongRepository) Get(group string, song string) (*song.Song, error) {
	panic("unimplemented")
}

func (s *SongRepository) Update(song *song.Song) error {
	panic("unimplemented")
}
