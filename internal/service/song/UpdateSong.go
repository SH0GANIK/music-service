package song

import (
	"context"
	"music-service/internal/model"
	"music-service/internal/repository/convertor"
)

func (r *serv) UpdateSong(ctx context.Context, song *model.Song) error {
	err := r.repo.UpdateSong(ctx, convertor.ToRepositorySongFromService(song))
	if err != nil {
		return err
	}
	return nil
}
