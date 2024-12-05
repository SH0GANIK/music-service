package convertor

import (
	"music-service/internal/logger"
	"music-service/internal/model"
	repository "music-service/internal/repository/model"
	"time"
)

func ToServiceSongFromRepository(song *repository.Song) *model.Song {
	return &model.Song{
		ID:    song.ID,
		Group: song.Group,
		Song:  song.Song,
		SongDetails: model.SongDetails{
			ReleaseDate: song.ReleaseDate.Format("02.01.2006"),
			Text:        song.Text,
			Link:        song.Link,
		},
	}
}

func ToRepositorySongFromService(song *model.Song) *repository.Song {
	t, err := time.Parse("02.01.2006", song.ReleaseDate)
	logger.Log.Debug(err, song.ReleaseDate)
	if song.ReleaseDate == "" {
		t = time.Time{}
	}
	return &repository.Song{
		ID:    song.ID,
		Group: song.Group,
		Song:  song.Song,
		SongDetails: repository.SongDetails{
			ReleaseDate: t,
			Text:        song.Text,
			Link:        song.Link,
		},
	}
}
