package mock

import "github.com/mishakrpv/musiclib/internal/infrastructure/service"

type GoodMusicInfoClient struct {
}

// GetSongDetail implements clients.MusicInfoClient.
func (c *GoodMusicInfoClient) GetSongDetail(group string, song string) (*service.SongDetail, error) {
	return &service.SongDetail{
		ReleaseDate: "16.07.2006",
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}, nil
}

func NewGoodMusicInfoClient() service.MusicInfoClient {
	return &GoodMusicInfoClient{}
}
