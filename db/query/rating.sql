-- name: CreateRating :one
INSERT INTO ratings (
    movie_id, user_id, rating_score
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetRating :one
SELECT * FROM ratings
WHERE movie_id = $1 AND user_id = $2
LIMIT 1;

-- name: UpdateRating :one
UPDATE ratings
SET 
    rating_score = COALESCE($3, rating_score),
    last_updated = CURRENT_TIMESTAMP
WHERE movie_id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteRating :exec
DELETE FROM ratings
WHERE movie_id = $1 AND user_id = $2;