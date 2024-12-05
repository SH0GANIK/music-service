package song

import (
	"github.com/jackc/pgx/v5"
	"music-service/internal/repository"
)

type repo struct {
	db *pgx.Conn
}

func NewSongRepository(db *pgx.Conn) repository.SongRepository {
	return &repo{
		db: db,
	}
}
