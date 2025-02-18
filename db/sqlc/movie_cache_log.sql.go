// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: movie_cache_log.sql

package db

import (
	"context"
)

const deleteMovieCache = `-- name: DeleteMovieCache :exec
DELETE FROM movie_cache_log
WHERE tmdb_id = $1
`

func (q *Queries) DeleteMovieCache(ctx context.Context, tmdbID int32) error {
	_, err := q.db.Exec(ctx, deleteMovieCache, tmdbID)
	return err
}

const getMovieCache = `-- name: GetMovieCache :one
SELECT tmdb_id, last_fetched FROM movie_cache_log
WHERE tmdb_id = $1 LIMIT 1
`

func (q *Queries) GetMovieCache(ctx context.Context, tmdbID int32) (MovieCacheLog, error) {
	row := q.db.QueryRow(ctx, getMovieCache, tmdbID)
	var i MovieCacheLog
	err := row.Scan(&i.TmdbID, &i.LastFetched)
	return i, err
}

const logMovieCache = `-- name: LogMovieCache :one
INSERT INTO movie_cache_log (
    tmdb_id
) VALUES (
    $1
)
RETURNING tmdb_id, last_fetched
`

func (q *Queries) LogMovieCache(ctx context.Context, tmdbID int32) (MovieCacheLog, error) {
	row := q.db.QueryRow(ctx, logMovieCache, tmdbID)
	var i MovieCacheLog
	err := row.Scan(&i.TmdbID, &i.LastFetched)
	return i, err
}

const updateMovieCache = `-- name: UpdateMovieCache :one
UPDATE movie_cache_log
SET 
    last_fetched = CURRENT_TIMESTAMP
WHERE tmdb_id = $1 
RETURNING tmdb_id, last_fetched
`

func (q *Queries) UpdateMovieCache(ctx context.Context, tmdbID int32) (MovieCacheLog, error) {
	row := q.db.QueryRow(ctx, updateMovieCache, tmdbID)
	var i MovieCacheLog
	err := row.Scan(&i.TmdbID, &i.LastFetched)
	return i, err
}
