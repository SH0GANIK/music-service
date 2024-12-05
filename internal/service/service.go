package service

import (
	"context"
	"music-service/internal/model"
)

type SongService interface {
	CreateSong(ctx context.Context, info *model.Song) (string, error)
	GetSongs(ctx context.Context, page, pageSize int, filters map[string]string) (*model.PaginatedSongsResponse, error)
	GetSongText(ctx context.Context, page, pageSize int, id string) (*model.PaginatedSongTextResponse, error)
	DeleteSong(ctx context.Context, id string) error
	UpdateSong(ctx context.Context, song *model.Song) error
}
