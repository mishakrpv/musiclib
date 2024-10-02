package song

type Repository interface {
	Create(song *Song) error
	Get(group string, song string) (*Song, error)
	FindMatching(predicate *Song) ([]*Song, error)
	Update(song *Song) error
	Delete(group string, song string) error
}
