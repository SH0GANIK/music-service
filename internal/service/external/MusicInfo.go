package external

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"music-service/internal/logger"
	"music-service/internal/model"
	"net/http"
	"net/url"
)

var (
	ErrBadRequest  = errors.New("bad request")
	ErrNotResponce = errors.New("not responce")
)

type MusicInfo struct {
	client *http.Client
	Url    string `json:"url"`
}

func NewMusicInfo(url string) *MusicInfo {
	return &MusicInfo{
		client: http.DefaultClient,
		Url:    url,
	}
}

func (m *MusicInfo) GetSongDetails(ctx context.Context, song *model.Song) (*model.SongDetails, error) {
	endpoint := "/info"
	params := url.Values{}
	params.Add("group", song.Group)
	params.Add("song", song.Song)
	urlWithParams := m.Url + endpoint + "?" + params.Encode()
	logger.Log.Debug("Request to external API sent. URL:" + urlWithParams)
	resp, err := m.client.Get(urlWithParams)
	if err != nil {
		return nil, fmt.Errorf("failed get response from MusicInfo: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return nil, ErrBadRequest
		} else {
			return nil, ErrNotResponce
		}
	}
	logger.Log.Debug("Response from external API received", resp.Body)
	var songDetails model.SongDetails
	if err := json.NewDecoder(resp.Body).Decode(&songDetails); err != nil {
		return nil, fmt.Errorf("failed to decode response from MusicInfo: %w", err)
	}
	return &songDetails, nil
}
