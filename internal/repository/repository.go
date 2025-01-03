package repository

import (
	"context"
	"music-service/internal/model"
	repository "music-service/internal/repository/model"
)

type SongRepository interface {
	CreateSong(ctx context.Context, info *repository.Song) (string, error)
	GetSongs(ctx context.Context, page, pageSize int, filters map[string]string) ([]*model.Song, error)
	GetSongText(ctx context.Context, id string) (string, error)
	UpdateSong(ctx context.Context, song *repository.Song) error
	DeleteSong(ctx context.Context, id string) error
}
