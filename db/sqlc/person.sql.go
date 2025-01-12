// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: person.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	decimal "github.com/shopspring/decimal"
)

const createPerson = `-- name: CreatePerson :one
INSERT INTO people (
    tmdb_id,
    name,
    known_for_department,
    biography,
    birthday,
    deathday,
    gender,
    profile_path,
    tmdb_popularity
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING person_id, tmdb_id, name, known_for_department, biography, birthday, deathday, gender, profile_path, tmdb_popularity, last_updated
`

type CreatePersonParams struct {
	TmdbID             int32           `json:"tmdb_id"`
	Name               string          `json:"name"`
	KnownForDepartment pgtype.Text     `json:"known_for_department"`
	Biography          pgtype.Text     `json:"biography"`
	Birthday           pgtype.Date     `json:"birthday"`
	Deathday           pgtype.Date     `json:"deathday"`
	Gender             pgtype.Int2     `json:"gender"`
	ProfilePath        pgtype.Text     `json:"profile_path"`
	TmdbPopularity     decimal.Decimal `json:"tmdb_popularity"`
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (People, error) {
	row := q.db.QueryRow(ctx, createPerson,
		arg.TmdbID,
		arg.Name,
		arg.KnownForDepartment,
		arg.Biography,
		arg.Birthday,
		arg.Deathday,
		arg.Gender,
		arg.ProfilePath,
		arg.TmdbPopularity,
	)
	var i People
	err := row.Scan(
		&i.PersonID,
		&i.TmdbID,
		&i.Name,
		&i.KnownForDepartment,
		&i.Biography,
		&i.Birthday,
		&i.Deathday,
		&i.Gender,
		&i.ProfilePath,
		&i.TmdbPopularity,
		&i.LastUpdated,
	)
	return i, err
}

const deletePerson = `-- name: DeletePerson :exec
DELETE FROM people
WHERE person_id = $1
`

func (q *Queries) DeletePerson(ctx context.Context, personID int32) error {
	_, err := q.db.Exec(ctx, deletePerson, personID)
	return err
}

const getPerson = `-- name: GetPerson :one
SELECT person_id, tmdb_id, name, known_for_department, biography, birthday, deathday, gender, profile_path, tmdb_popularity, last_updated FROM people
WHERE person_id = $1 LIMIT 1
`

func (q *Queries) GetPerson(ctx context.Context, personID int32) (People, error) {
	row := q.db.QueryRow(ctx, getPerson, personID)
	var i People
	err := row.Scan(
		&i.PersonID,
		&i.TmdbID,
		&i.Name,
		&i.KnownForDepartment,
		&i.Biography,
		&i.Birthday,
		&i.Deathday,
		&i.Gender,
		&i.ProfilePath,
		&i.TmdbPopularity,
		&i.LastUpdated,
	)
	return i, err
}

const updatePerson = `-- name: UpdatePerson :one
UPDATE people
SET 
    tmdb_id = COALESCE($2, tmdb_id), 
    name = COALESCE($3, name), 
    known_for_department = COALESCE($4, known_for_department), 
    biography = COALESCE($5, biography), 
    birthday = COALESCE($6, birthday), 
    deathday = COALESCE($7, deathday),
    gender = COALESCE($8, gender),
    profile_path = COALESCE($9, profile_path),
    tmdb_popularity = COALESCE($10, tmdb_popularity),
    last_updated = CURRENT_TIMESTAMP
WHERE person_id = $1
RETURNING person_id, tmdb_id, name, known_for_department, biography, birthday, deathday, gender, profile_path, tmdb_popularity, last_updated
`

type UpdatePersonParams struct {
	PersonID           int32           `json:"person_id"`
	TmdbID             int32           `json:"tmdb_id"`
	Name               string          `json:"name"`
	KnownForDepartment pgtype.Text     `json:"known_for_department"`
	Biography          pgtype.Text     `json:"biography"`
	Birthday           pgtype.Date     `json:"birthday"`
	Deathday           pgtype.Date     `json:"deathday"`
	Gender             pgtype.Int2     `json:"gender"`
	ProfilePath        pgtype.Text     `json:"profile_path"`
	TmdbPopularity     decimal.Decimal `json:"tmdb_popularity"`
}

func (q *Queries) UpdatePerson(ctx context.Context, arg UpdatePersonParams) (People, error) {
	row := q.db.QueryRow(ctx, updatePerson,
		arg.PersonID,
		arg.TmdbID,
		arg.Name,
		arg.KnownForDepartment,
		arg.Biography,
		arg.Birthday,
		arg.Deathday,
		arg.Gender,
		arg.ProfilePath,
		arg.TmdbPopularity,
	)
	var i People
	err := row.Scan(
		&i.PersonID,
		&i.TmdbID,
		&i.Name,
		&i.KnownForDepartment,
		&i.Biography,
		&i.Birthday,
		&i.Deathday,
		&i.Gender,
		&i.ProfilePath,
		&i.TmdbPopularity,
		&i.LastUpdated,
	)
	return i, err
}
