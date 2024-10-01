package song

type Repository interface {
	CreateSong(song *Song) error
	GetSong(group string, song string) (*Song, error)
	GetAllSongs() ([]Song, error)
	UpdateSong(song *Song) error
	DeleteSong(group string, song string) error
}
