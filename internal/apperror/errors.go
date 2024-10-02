package apperror

import "errors"

var (
	ErrVerseNotFound = errors.New("verse not found")
	ErrSongNotFound  = errors.New("song not found")
)
