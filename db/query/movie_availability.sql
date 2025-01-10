-- name: CreateMovieAvailability :one
INSERT INTO movie_availability (
    movie_id, 
    service_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetMovieAvailabilities :many
SELECT * FROM movie_availability
WHERE movie_id = $1
ORDER BY service_id
LIMIT $2
OFFSET $3;

-- name: DeleteMovieAvailability :exec
DELETE FROM movie_availability
WHERE movie_id = $1 AND service_id = $2;