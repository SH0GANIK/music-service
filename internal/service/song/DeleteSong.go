package song

import (
	"context"
)

func (r *serv) DeleteSong(ctx context.Context, id string) error {
	err := r.repo.DeleteSong(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
