package song

type Repository interface {
	Create(song *Song) error
	Get(id string) (*Song, error)
	FindMatching(predicate *Song) ([]*Song, error)
	Update(song *Song) error
	Delete(id string) error
}
