-- name: CreateComparison :one
INSERT INTO comparisons (
    user_id, base_movie_id, compared_movie_id, preference
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetComparison :one
SELECT * FROM comparisons
WHERE comparison_id = $1 LIMIT 1;

-- name: UpdateComparison :one
UPDATE comparisons
SET 
    preference = COALESCE($2, preference)
WHERE comparison_id = $1
RETURNING *;

-- name: DeleteComparison :exec
DELETE FROM comparisons
WHERE comparison_id = $1;