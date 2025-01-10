-- name: LogMovieCache :one
INSERT INTO movie_cache_log (
    tmdb_id
) VALUES (
    $1
)
RETURNING *;

-- name: GetMovieCache :one
SELECT * FROM movie_cache_log
WHERE tmdb_id = $1 LIMIT 1;

-- name: UpdateMovieCache :one
UPDATE movie_cache_log
SET 
    last_fetched = CURRENT_TIMESTAMP
WHERE tmdb_id = $1 
RETURNING *;

-- name: DeleteMovieCache :exec
DELETE FROM movie_cache_log
WHERE tmdb_id = $1;