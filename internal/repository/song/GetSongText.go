package song

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *repo) GetSongText(ctx context.Context, id string) (string, error) {
	query := `SELECT text FROM songs WHERE id=$1`
	var text string
	if err := r.db.QueryRow(ctx, query, id).Scan(&text); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("song with id %s not found", id)
		}
		return "", fmt.Errorf("failed to get song text from database: %w", err)
	}
	return text, nil
}
