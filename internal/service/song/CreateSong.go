package song

import (
	"context"
	"fmt"
	"music-service/internal/model"
	"music-service/internal/repository/convertor"
)

func (r *serv) CreateSong(ctx context.Context, info *model.Song) (string, error) {
	songDetails, err := r.musicInfo.GetSongDetails(ctx, info)
	if err != nil {
		return "", fmt.Errorf("failed to get song details: %w", err)
	}
	info.SongDetails = *songDetails
	if err != nil {
		return "", fmt.Errorf("Error parsing date: %w", err)
	}
	id, err := r.repo.CreateSong(ctx, convertor.ToRepositorySongFromService(info))
	if err != nil {
		return "", err
	}
	return id, nil
}
