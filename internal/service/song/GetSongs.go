package song

import (
	"context"
	"fmt"
	"music-service/internal/model"
)

func (r *serv) GetSongs(ctx context.Context, page, pageSize int, filters map[string]string) (*model.PaginatedSongsResponse, error) {
	songs, err := r.repo.GetSongs(ctx, page, pageSize, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to get songs: %w", err)
	}
	return &model.PaginatedSongsResponse{Data: songs, Page: page, PageSize: pageSize}, nil
}
