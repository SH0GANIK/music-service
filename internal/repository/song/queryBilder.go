package song

import (
	repository "music-service/internal/repository/model"
	"strconv"
)

type QueryBuilder struct {
	query  string
	params []interface{}
}

func (qb *QueryBuilder) Append(sql string) *QueryBuilder {
	qb.query += " " + sql
	return qb
}

func (qb *QueryBuilder) AppendParams(params ...interface{}) *QueryBuilder {
	qb.params = append(qb.params, params...)
	return qb
}

func (qb *QueryBuilder) AppendFilters(filters map[string]string) *QueryBuilder {
	if len(filters) == 0 {
		return qb
	}
	qb.query += " WHERE 1=1"
	i := 1
	for key, value := range filters {
		switch key {
		case "group":
			qb.query += " AND group_name = $" + strconv.Itoa(i)
			qb.params = append(qb.params, value)
		case "song":
			qb.query += " AND song LIKE $" + strconv.Itoa(i)
			qb.params = append(qb.params, value)
		case "releaseDate":
			qb.query += " AND release_date >= $" + strconv.Itoa(i)
			qb.params = append(qb.params, value)
		}
		i++
	}
	return qb
}

func (qb *QueryBuilder) UpdateSongQuery(song *repository.Song) *QueryBuilder {
	qb.query += "UPDATE songs SET"
	i := 1
	if song.Group != "" {
		qb.Append("group_name = $" + strconv.Itoa(i))
		qb.AppendParams(song.Group)
		i++
	}
	if song.Song != "" {
		qb.Append("song = $" + strconv.Itoa(i))
		qb.AppendParams(song.Song)
		i++
	}
	if !song.ReleaseDate.IsZero() {
		qb.Append("release_date = $" + strconv.Itoa(i))
		qb.AppendParams(song.ReleaseDate)
		i++
	}
	if song.Text != "" {
		qb.Append("text = $" + strconv.Itoa(i))
		qb.AppendParams(song.Text)
		i++
	}
	if song.Link != "" {
		qb.Append("link = $" + strconv.Itoa(i))
		qb.AppendParams(song.Text)
		i++
	}
	qb.Append("WHERE id = $" + strconv.Itoa(i))
	qb.AppendParams(song.ID)
	return qb
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	return qb.query, qb.params
}
