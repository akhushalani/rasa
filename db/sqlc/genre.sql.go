// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: genre.sql

package db

import (
	"context"
)

const createGenre = `-- name: CreateGenre :one
INSERT INTO genres (
    name
) VALUES (
    $1
)
RETURNING genre_id, name
`

func (q *Queries) CreateGenre(ctx context.Context, name string) (Genres, error) {
	row := q.db.QueryRow(ctx, createGenre, name)
	var i Genres
	err := row.Scan(&i.GenreID, &i.Name)
	return i, err
}

const deleteGenre = `-- name: DeleteGenre :exec
DELETE FROM genres
WHERE genre_id = $1
`

func (q *Queries) DeleteGenre(ctx context.Context, genreID int32) error {
	_, err := q.db.Exec(ctx, deleteGenre, genreID)
	return err
}

const getGenre = `-- name: GetGenre :one
SELECT genre_id, name FROM genres
WHERE genre_id = $1 LIMIT 1
`

func (q *Queries) GetGenre(ctx context.Context, genreID int32) (Genres, error) {
	row := q.db.QueryRow(ctx, getGenre, genreID)
	var i Genres
	err := row.Scan(&i.GenreID, &i.Name)
	return i, err
}

const updateGenre = `-- name: UpdateGenre :one
UPDATE genres
SET 
    name = COALESCE($2, name)
WHERE genre_id = $1
RETURNING genre_id, name
`

type UpdateGenreParams struct {
	GenreID int32  `json:"genre_id"`
	Name    string `json:"name"`
}

func (q *Queries) UpdateGenre(ctx context.Context, arg UpdateGenreParams) (Genres, error) {
	row := q.db.QueryRow(ctx, updateGenre, arg.GenreID, arg.Name)
	var i Genres
	err := row.Scan(&i.GenreID, &i.Name)
	return i, err
}
