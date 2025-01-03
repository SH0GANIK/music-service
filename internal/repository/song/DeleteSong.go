package song

import (
	"context"
	"fmt"
)

func (r *repo) DeleteSong(ctx context.Context, id string) error {
	queryDelete := `DELETE FROM songs WHERE id = $1`
	res, err := r.db.Exec(ctx, queryDelete, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("song with id %s not found", id)
	}
	return nil
}
