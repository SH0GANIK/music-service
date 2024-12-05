package song

import (
	"context"
	"fmt"
	"music-service/internal/model"
	"strings"
)

func (r *serv) GetSongText(ctx context.Context, page, pageSize int, id string) (*model.PaginatedSongTextResponse, error) {
	text, err := r.repo.GetSongText(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get song text: %w", err)
	}
	data := strings.Split(text, "\n")
	startIndex := (page - 1) * pageSize
	endIndex := min(len(data), startIndex+pageSize)
	return &model.PaginatedSongTextResponse{Data: data[startIndex:endIndex], Page: page, PageSize: pageSize}, nil
}
