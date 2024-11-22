package musicinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mishakrpv/musiclib/internal/apperror"
	"github.com/rs/zerolog/log"
)

type Client interface {
	GetSongDetail(group string, song string) (*SongDetail, error)
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type HTTPClient struct {
	serviceBaseUrl string
}

func NewHTTPMusicInfoClient(serviceBaseUrl string) Client {
	return &HTTPClient{serviceBaseUrl: serviceBaseUrl}
}

func (h *HTTPClient) GetSongDetail(group string, song string) (*SongDetail, error) {
	log.Info().Str("service_base_url", h.serviceBaseUrl).
		Msg("Creating http request")

	url := fmt.Sprintf("%s/info?group=%s&song=%s", h.serviceBaseUrl, group, song)
	url = strings.Replace(url, " ", "+", -1)

	log.Debug().Str("url", url).Msg("Info url")

	request, err := http.NewRequest(http.MethodGet,
		url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error creating request")
		return nil, err
	}

	log.Info().Msg("Request created")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("Error sending request")
		return nil, err
	}
	defer response.Body.Close()

	log.Info().
		Msg("Request sent")

	if response.StatusCode == http.StatusNotFound {
		log.Warn().Msg("Response does not indicate success")
		return nil, apperror.ErrSongNotFound
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error reading request")
		return nil, err
	}

	log.Info().
		Msg("Response body read")

	detail := &SongDetail{}

	err = json.Unmarshal(body, &detail)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshaling response body")
		return nil, err
	}

	log.Info().
		Msg("Response body unmarshaled")

	return detail, nil
}
