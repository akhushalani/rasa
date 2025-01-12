-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password_hash
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT 
    user_id, 
    name, 
    email, 
    password_hash, 
    created_at 
FROM users 
WHERE user_id = $1;

-- name: UpdateUser :exec
UPDATE users
SET 
    name = COALESCE($2, name),
    email = COALESCE($3, email),
    password_hash = COALESCE($4, password_hash)
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE user_id = $1;
