package song

import (
	"context"
	"fmt"
	repository "music-service/internal/repository/model"
)

func (r *repo) CreateSong(ctx context.Context, info *repository.Song) (string, error) {
	queryExist := `SELECT EXISTS (SELECT 1 FROM songs WHERE group_name = $1 AND song = $2);`
	exist := false
	if err := r.db.QueryRow(ctx, queryExist, info.Group, info.Song).Scan(&exist); err != nil {
		return "", fmt.Errorf("failed to check if song exists: %w", err)
	}
	if exist {
		return "", fmt.Errorf("song with group '%s' and name '%s' already exists", info.Group, info.Song)
	}
	queryInsert := `INSERT INTO songs
		(group_name, song, release_date, text, link)
		VALUES
		($1, $2, $3, $4, $5)
		RETURNING id`
	var id string
	err := r.db.QueryRow(ctx, queryInsert, info.Group, info.Song, info.ReleaseDate, info.Text, info.Link).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to create song in database: %w", err)
	}
	return id, nil
}
