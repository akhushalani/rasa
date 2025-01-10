-- name: CreateUserMovie :one
INSERT INTO user_movies (
    user_id, movie_id, rating, review, watchlist, watched, favorited
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetUserMovie :one
SELECT * FROM user_movies
WHERE user_id = $1 AND movie_id = $2
LIMIT 1;

-- name: ListUserMovies :many
SELECT * FROM user_movies
WHERE user_id = $1
ORDER BY movie_id
LIMIT $2
OFFSET $3;

-- name: UpdateUserMovie :one
UPDATE user_movies
SET 
    rating = COALESCE($3, rating), 
    review = COALESCE($4, review), 
    watchlist = COALESCE($5, watchlist), 
    watched = COALESCE($6, watched), 
    favorited = COALESCE($7, favorited)
WHERE user_id = $1 AND movie_id = $2
RETURNING *;