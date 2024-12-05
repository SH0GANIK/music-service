package song

import (
	"context"
	"fmt"
	"music-service/internal/logger"
	repository "music-service/internal/repository/model"
)

func (r *repo) UpdateSong(ctx context.Context, song *repository.Song) error {
	qb := &QueryBuilder{}
	qb.UpdateSongQuery(song)
	query, params := qb.Build()
	logger.Log.Debug("UpdateSong Query:", query, "Params:", params)
	res, err := r.db.Exec(ctx, query, params...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("song with id %s not found", song.ID)
	}
	return nil
}
