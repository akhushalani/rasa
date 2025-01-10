-- name: CreateGenre :one
INSERT INTO genres (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: GetGenre :one
SELECT * FROM genres
WHERE genre_id = $1 LIMIT 1;

-- name: UpdateGenre :one
UPDATE genres
SET 
    name = COALESCE($2, name)
WHERE genre_id = $1
RETURNING *;

-- name: DeleteGenre :exec
DELETE FROM genres
WHERE genre_id = $1;