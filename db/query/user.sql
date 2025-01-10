-- name: CreateUser :exec
INSERT INTO users (
    username,
    email,
    password_hash
) VALUES (
    $1, $2, $3
);

-- name: GetUserByID :one
SELECT 
    user_id, 
    username, 
    email, 
    password_hash, 
    created_at 
FROM users 
WHERE user_id = $1;

-- name: GetUserByUsername :one
SELECT 
    user_id, 
    username, 
    email, 
    password_hash, 
    created_at 
FROM users 
WHERE username = $1;

-- name: UpdateUser :exec
UPDATE users
SET 
    username = COALESCE($2, username),
    email = COALESCE($3, email),
    password_hash = COALESCE($4, password_hash)
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE user_id = $1;
