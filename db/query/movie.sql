-- name: CreateMovie :one
INSERT INTO movies (
    tmdb_id, imdb_id, title, overview, release_date, poster_path, backdrop_path, tmdb_popularity
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetMovie :one
SELECT * FROM movies
WHERE movie_id = $1 LIMIT 1;

-- name: GetMovieByTmdbId :one
SELECT * FROM movies
WHERE tmdb_id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies
ORDER BY movie_id
LIMIT $1
OFFSET $2;

-- name: UpdateMovie :one
UPDATE movies
SET 
    tmdb_id = COALESCE($2, tmdb_id), 
    imdb_id = COALESCE($3, imdb_id), 
    title = COALESCE($4, title), 
    overview = COALESCE($5, overview), 
    release_date = COALESCE($6, release_date), 
    poster_path = COALESCE($7, poster_path), 
    backdrop_path = COALESCE($8, backdrop_path), 
    tmdb_popularity = COALESCE($9, tmdb_popularity), 
    last_updated = CURRENT_TIMESTAMP
WHERE movie_id = $1
RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM movies
WHERE movie_id = $1;