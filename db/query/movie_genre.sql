-- name: CreateMovieGenre :one
INSERT INTO movie_genres (
    movie_id, genre_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetMovieGenre :one
SELECT * FROM movie_genres
WHERE movie_id = $1 AND genre_id = $2
LIMIT 1;

-- name: ListMovieGenres :many
SELECT * FROM movie_genres
ORDER BY movie_id
LIMIT $1
OFFSET $2;

-- name: DeleteMovieGenre :exec
DELETE FROM movie_genres
WHERE movie_id = $1 AND genre_id = $2;