package song

import (
	"context"
	"fmt"
	"music-service/internal/logger"
	"music-service/internal/model"
	"music-service/internal/repository/convertor"
	repository "music-service/internal/repository/model"
	"strconv"
)

func (r *repo) GetSongs(ctx context.Context, page, pageSize int, filters map[string]string) ([]*model.Song, error) {
	qb := &QueryBuilder{}
	qb.Append("SELECT id, group_name, song, release_date, text, link FROM songs")
	qb.AppendFilters(filters)
	qb.Append("ORDER BY release_date DESC LIMIT $" + strconv.Itoa(len(filters)+1) + " OFFSET $" + strconv.Itoa(len(filters)+2))
	qb.AppendParams(pageSize, (page-1)*pageSize)
	query, params := qb.Build()
	rows, err := r.db.Query(ctx, query, params...)
	logger.Log.Debug("GetSongs Query: "+query, "Params: ", params)
	if err != nil {
		return nil, fmt.Errorf("error getting songs from database: %w", err)
	}
	defer rows.Close()
	var songs []*model.Song
	for rows.Next() {
		song := &repository.Song{}
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			return nil, fmt.Errorf("error scanning song from database: %w", err)
		}
		songs = append(songs, convertor.ToServiceSongFromRepository(song))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error from rows: %w", err)
	}
	return songs, nil
}
