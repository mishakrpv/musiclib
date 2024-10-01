package clients

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type MusicInfoClient interface {
	GetSongDetail(group string, song string) (*SongDetail, error)
}

type SongDetail struct {
	ReleaseDate time.Time
	Text        string
	Link        string
}

type HttpMusicInfoClient struct {
	serviceBaseUrl string
}

func NewHttpMusicInfoClient(serviceBaseUrl string) MusicInfoClient {
	return &HttpMusicInfoClient{serviceBaseUrl: serviceBaseUrl}
}

func (h *HttpMusicInfoClient) GetSongDetail(group string, song string) (*SongDetail, error) {
	zap.L().Debug("Creating http request", zap.String("service_base_url", h.serviceBaseUrl))

	request, err := http.NewRequest(http.MethodGet, h.serviceBaseUrl+"/info", nil)
	if err != nil {
		zap.L().Error("Error creating request", zap.Error(err))
		return nil, err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		zap.L().Error("Error sending request", zap.Error(err))
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		zap.L().Error("Error reading request", zap.Error(err))
		return nil, err
	}

	detail := &SongDetail{}

	err = json.Unmarshal(body, &detail)
	if err != nil {
		zap.L().Error("Error unmarshaling response body", zap.Error(err))
		return nil, err
	}

	return detail, nil
}
