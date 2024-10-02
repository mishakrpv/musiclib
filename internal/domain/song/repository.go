package song

type Repository interface {
	Create(song *Song) error
	Get(groupName string, songName string) (*Song, error)
	FindMatching(predicate *Song) ([]*Song, error)
	Update(song *Song) error
	Delete(groupName string, songName string) error
}
