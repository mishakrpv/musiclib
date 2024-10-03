package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mishakrpv/musiclib/internal/apperror"
	"go.uber.org/zap"
)

type MusicInfoClient interface {
	GetSongDetail(group string, song string) (*SongDetail, error)
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type HttpMusicInfoClient struct {
	serviceBaseUrl string
}

func NewHttpMusicInfoClient(serviceBaseUrl string) MusicInfoClient {
	return &HttpMusicInfoClient{serviceBaseUrl: serviceBaseUrl}
}

func (h *HttpMusicInfoClient) GetSongDetail(group string, song string) (*SongDetail, error) {
	zap.L().Info("Creating http request", zap.String("service_base_url", h.serviceBaseUrl))

	url := fmt.Sprintf("%s/info?group=%s&song=%s", h.serviceBaseUrl, group, song)
	url = strings.Replace(url, " ", "+", -1)

	zap.L().Debug("Info url", zap.String("url", url))

	request, err := http.NewRequest(http.MethodGet,
		url, nil)
	if err != nil {
		zap.L().Error("Error creating request", zap.Error(err))
		return nil, err
	}

	zap.L().Info("Request created")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		zap.L().Error("Error sending request", zap.Error(err))
		return nil, err
	}
	defer response.Body.Close()

	zap.L().Info("Request sent")

	if response.StatusCode == http.StatusNotFound {
		zap.L().Warn("Response does not indicate success")
		return nil, apperror.ErrSongNotFound
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		zap.L().Error("Error reading request", zap.Error(err))
		return nil, err
	}

	zap.L().Info("Response body read")

	detail := &SongDetail{}

	err = json.Unmarshal(body, &detail)
	if err != nil {
		zap.L().Error("Error unmarshaling response body", zap.Error(err))
		return nil, err
	}

	zap.L().Info("Response body unmarshaled")

	return detail, nil
}
