package song

import (
	"music-service/internal/repository"
	"music-service/internal/service"
	"music-service/internal/service/external"
)

type serv struct {
	repo      repository.SongRepository
	musicInfo *external.MusicInfo
}

func NewSongService(repo repository.SongRepository, musicInfo *external.MusicInfo) service.SongService {
	return &serv{
		repo:      repo,
		musicInfo: musicInfo,
	}
}
